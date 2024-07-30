//go:build integration

package webkey_test

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/zitadel/zitadel/internal/integration"
	"github.com/zitadel/zitadel/pkg/grpc/feature/v2"
	webkey "github.com/zitadel/zitadel/pkg/grpc/webkey/v3alpha"
)

var (
	CTX       context.Context
	SystemCTX context.Context
	Tester    *integration.Tester
)

func TestMain(m *testing.M) {
	os.Exit(func() int {
		ctx, _, cancel := integration.Contexts(time.Hour)
		defer cancel()

		Tester = integration.NewTester(ctx)
		defer Tester.Done()

		SystemCTX = Tester.WithAuthorization(ctx, integration.SystemUser)
		CTX = Tester.WithAuthorization(ctx, integration.IAMOwner)
		return m.Run()
	}())
}

func TestServer_Feature_Disabled(t *testing.T) {
	client, iamCTX := createInstanceAndClients(t, false)

	t.Run("GenerateWebKey", func(t *testing.T) {
		_, err := client.GenerateWebKey(iamCTX, &webkey.GenerateWebKeyRequest{})
		assertFeatureDisabledError(t, err)
	})
	t.Run("ActivateWebKey", func(t *testing.T) {
		_, err := client.ActivateWebKey(iamCTX, &webkey.ActivateWebKeyRequest{
			KeyId: "1",
		})
		assertFeatureDisabledError(t, err)
	})
	t.Run("DeleteWebKey", func(t *testing.T) {
		_, err := client.DeleteWebKey(iamCTX, &webkey.DeleteWebKeyRequest{
			KeyId: "1",
		})
		assertFeatureDisabledError(t, err)
	})
	t.Run("ListWebKeys", func(t *testing.T) {
		_, err := client.ListWebKeys(iamCTX, &webkey.ListWebKeysRequest{})
		assertFeatureDisabledError(t, err)
	})
}

func TestServer_ListWebKeys(t *testing.T) {
	client, iamCtx := createInstanceAndClients(t, true)
	// After the feature is first enabled, we can expect 2 generated keys with the default config.
	checkWebKeyListState(iamCtx, t, client, 2, "", &webkey.WebKeyDetails_Rsa{
		Rsa: &webkey.WebKeyRSAConfig{
			Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
			Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
		},
	})
}

func TestServer_GenerateWebKey(t *testing.T) {
	client, iamCtx := createInstanceAndClients(t, true)
	_, err := client.GenerateWebKey(iamCtx, &webkey.GenerateWebKeyRequest{
		Config: &webkey.GenerateWebKeyRequest_Rsa{
			Rsa: &webkey.WebKeyRSAConfig{
				Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
				Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
			},
		},
	})
	require.NoError(t, err)

	checkWebKeyListState(iamCtx, t, client, 3, "", &webkey.WebKeyDetails_Rsa{
		Rsa: &webkey.WebKeyRSAConfig{
			Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
			Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
		},
	})
}

func TestServer_ActivateWebKey(t *testing.T) {
	client, iamCtx := createInstanceAndClients(t, true)
	resp, err := client.GenerateWebKey(iamCtx, &webkey.GenerateWebKeyRequest{
		Config: &webkey.GenerateWebKeyRequest_Rsa{
			Rsa: &webkey.WebKeyRSAConfig{
				Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
				Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
			},
		},
	})
	require.NoError(t, err)

	_, err = client.ActivateWebKey(iamCtx, &webkey.ActivateWebKeyRequest{
		KeyId: resp.GetKeyId(),
	})
	require.NoError(t, err)

	checkWebKeyListState(iamCtx, t, client, 3, resp.GetKeyId(), &webkey.WebKeyDetails_Rsa{
		Rsa: &webkey.WebKeyRSAConfig{
			Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
			Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
		},
	})
}

func TestServer_DeleteWebKey(t *testing.T) {
	client, iamCtx := createInstanceAndClients(t, true)
	keyIDs := make([]string, 2)
	for i := 0; i < 2; i++ {
		resp, err := client.GenerateWebKey(iamCtx, &webkey.GenerateWebKeyRequest{
			Config: &webkey.GenerateWebKeyRequest_Rsa{
				Rsa: &webkey.WebKeyRSAConfig{
					Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
					Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
				},
			},
		})
		require.NoError(t, err)
		keyIDs[i] = resp.GetKeyId()
	}
	_, err := client.ActivateWebKey(iamCtx, &webkey.ActivateWebKeyRequest{
		KeyId: keyIDs[0],
	})
	require.NoError(t, err)

	ok := t.Run("cannot delete active key", func(t *testing.T) {
		_, err := client.DeleteWebKey(iamCtx, &webkey.DeleteWebKeyRequest{
			KeyId: keyIDs[0],
		})
		require.Error(t, err)
		s := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, s.Code())
		assert.Contains(t, s.Message(), "COMMAND-Chai1")
	})
	if !ok {
		return
	}

	ok = t.Run("delete inactive key", func(t *testing.T) {
		_, err := client.DeleteWebKey(iamCtx, &webkey.DeleteWebKeyRequest{
			KeyId: keyIDs[1],
		})
		require.NoError(t, err)
	})
	if !ok {
		return
	}

	// There are 2 keys from feature setup, +2 created, -1 deleted = 3
	checkWebKeyListState(iamCtx, t, client, 3, keyIDs[0], &webkey.WebKeyDetails_Rsa{
		Rsa: &webkey.WebKeyRSAConfig{
			Bits:   webkey.WebKeyRSAConfig_RSA_BITS_2048,
			Hasher: webkey.WebKeyRSAConfig_RSA_HASHER_SHA256,
		},
	})
}

func createInstanceAndClients(t *testing.T, enableFeature bool) (webkey.WebKeyServiceClient, context.Context) {
	domain, _, _, iamCTX := Tester.UseIsolatedInstance(t, CTX, SystemCTX)
	cc, err := grpc.NewClient(
		net.JoinHostPort(domain, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	if enableFeature {
		features := feature.NewFeatureServiceClient(cc)
		_, err = features.SetInstanceFeatures(iamCTX, &feature.SetInstanceFeaturesRequest{
			WebKey: proto.Bool(true),
		})
		require.NoError(t, err)
		time.Sleep(time.Second)
	}

	return webkey.NewWebKeyServiceClient(cc), iamCTX
}

func assertFeatureDisabledError(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
	s := status.Convert(err)
	assert.Equal(t, codes.FailedPrecondition, s.Code())
	assert.Contains(t, s.Message(), "WEBKEY-Ohx6E")
}

func checkWebKeyListState(ctx context.Context, t *testing.T, client webkey.WebKeyServiceClient, nKeys int, expectActiveKeyID string, config any) {
	resp, err := client.ListWebKeys(ctx, &webkey.ListWebKeysRequest{})
	require.NoError(t, err)
	list := resp.GetWebKeys()
	require.Len(t, list, nKeys)

	now := time.Now()
	var gotActiveKeyID string
	for _, key := range list {
		assert.NotEmpty(t, key.GetKeyId())
		assert.WithinRange(t, key.GetCreatedDate().AsTime(), now.Add(-time.Minute), now.Add(time.Minute))
		assert.WithinRange(t, key.GetChangeDate().AsTime(), now.Add(-time.Minute), now.Add(time.Minute))
		assert.NotEmpty(t, key.GetSequence())
		assert.NotEqual(t, webkey.WebKeyState_STATE_UNSPECIFIED, key.GetState())
		assert.NotEqual(t, webkey.WebKeyState_STATE_REMOVED, key.GetState())
		assert.Equal(t, config, key.GetConfig())

		if key.GetState() == webkey.WebKeyState_STATE_ACTIVE {
			gotActiveKeyID = key.GetKeyId()
		}
	}
	assert.NotEmpty(t, gotActiveKeyID)
	if expectActiveKeyID != "" {
		assert.Equal(t, expectActiveKeyID, gotActiveKeyID)
	}
}
