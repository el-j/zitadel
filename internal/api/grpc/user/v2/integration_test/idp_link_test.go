//go:build integration

package user_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zitadel/zitadel/internal/integration"
	"github.com/zitadel/zitadel/pkg/grpc/object/v2"
	"github.com/zitadel/zitadel/pkg/grpc/user/v2"
)

func TestServer_AddIDPLink(t *testing.T) {
	idpResp := Instance.AddGenericOAuthProvider(IamCTX, Instance.DefaultOrg.Id)
	type args struct {
		ctx context.Context
		req *user.AddIDPLinkRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.AddIDPLinkResponse
		wantErr bool
	}{
		{
			name: "user does not exist",
			args: args{
				CTX,
				&user.AddIDPLinkRequest{
					UserId: "userID",
					IdpLink: &user.IDPLink{
						IdpId:    idpResp.Id,
						UserId:   "userID",
						UserName: "username",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "idp does not exist",
			args: args{
				CTX,
				&user.AddIDPLinkRequest{
					UserId: Instance.Users.Get(integration.UserTypeOrgOwner).ID,
					IdpLink: &user.IDPLink{
						IdpId:    "idpID",
						UserId:   "userID",
						UserName: "username",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "add link",
			args: args{
				CTX,
				&user.AddIDPLinkRequest{
					UserId: Instance.Users.Get(integration.UserTypeOrgOwner).ID,
					IdpLink: &user.IDPLink{
						IdpId:    idpResp.Id,
						UserId:   "userID",
						UserName: "username",
					},
				},
			},
			want: &user.AddIDPLinkResponse{
				Details: &object.Details{
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Instance.DefaultOrg.Id,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.AddIDPLink(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			integration.AssertDetails(t, tt.want, got)
		})
	}
}

func TestServer_ListIDPLinks(t *testing.T) {
	orgResp := Instance.CreateOrganization(IamCTX, fmt.Sprintf("ListIDPLinks%d", time.Now().UnixNano()), fmt.Sprintf("%d@mouse.com", time.Now().UnixNano()))

	instanceIdpResp := Instance.AddGenericOAuthProvider(IamCTX, Instance.DefaultOrg.Id)
	userInstanceResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))
	_, err := Instance.CreateUserIDPlink(IamCTX, userInstanceResp.GetUserId(), "external_instance", instanceIdpResp.Id, "externalUsername_instance")
	require.NoError(t, err)
	orgIdpResp := Instance.AddOrgGenericOAuthProvider(IamCTX, orgResp.OrganizationId)
	userOrgResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))
	_, err = Instance.CreateUserIDPlink(IamCTX, userOrgResp.GetUserId(), "external_org", orgIdpResp.Id, "externalUsername_org")
	require.NoError(t, err)

	userMultipleResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))
	_, err = Instance.CreateUserIDPlink(IamCTX, userMultipleResp.GetUserId(), "external_multi", instanceIdpResp.Id, "externalUsername_multi")
	require.NoError(t, err)
	_, err = Instance.CreateUserIDPlink(IamCTX, userMultipleResp.GetUserId(), "external_multi", orgIdpResp.Id, "externalUsername_multi")
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		req *user.ListIDPLinksRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.ListIDPLinksResponse
		wantErr bool
	}{
		{
			name: "list links, no permission",
			args: args{
				UserCTX,
				&user.ListIDPLinksRequest{
					UserId: userOrgResp.GetUserId(),
				},
			},
			want: &user.ListIDPLinksResponse{
				Details: &object.ListDetails{
					TotalResult: 0,
					Timestamp:   timestamppb.Now(),
				},
				Result: []*user.IDPLink{},
			},
		},
		{
			name: "list links, no permission, org",
			args: args{
				CTX,
				&user.ListIDPLinksRequest{
					UserId: userOrgResp.GetUserId(),
				},
			},
			want: &user.ListIDPLinksResponse{
				Details: &object.ListDetails{
					TotalResult: 0,
					Timestamp:   timestamppb.Now(),
				},
				Result: []*user.IDPLink{},
			},
		},
		{
			name: "list idp links, org, ok",
			args: args{
				IamCTX,
				&user.ListIDPLinksRequest{
					UserId: userOrgResp.GetUserId(),
				},
			},
			want: &user.ListIDPLinksResponse{
				Details: &object.ListDetails{
					TotalResult: 1,
					Timestamp:   timestamppb.Now(),
				},
				Result: []*user.IDPLink{
					{
						IdpId:    orgIdpResp.Id,
						UserId:   "external_org",
						UserName: "externalUsername_org",
					},
				},
			},
		},
		{
			name: "list idp links, instance, ok",
			args: args{
				IamCTX,
				&user.ListIDPLinksRequest{
					UserId: userInstanceResp.GetUserId(),
				},
			},
			want: &user.ListIDPLinksResponse{
				Details: &object.ListDetails{
					TotalResult: 1,
					Timestamp:   timestamppb.Now(),
				},
				Result: []*user.IDPLink{
					{
						IdpId:    instanceIdpResp.Id,
						UserId:   "external_instance",
						UserName: "externalUsername_instance",
					},
				},
			},
		},
		{
			name: "list idp links, multi, ok",
			args: args{
				IamCTX,
				&user.ListIDPLinksRequest{
					UserId: userMultipleResp.GetUserId(),
				},
			},
			want: &user.ListIDPLinksResponse{
				Details: &object.ListDetails{
					TotalResult: 2,
					Timestamp:   timestamppb.Now(),
				},
				Result: []*user.IDPLink{
					{
						IdpId:    instanceIdpResp.Id,
						UserId:   "external_multi",
						UserName: "externalUsername_multi",
					},
					{
						IdpId:    orgIdpResp.Id,
						UserId:   "external_multi",
						UserName: "externalUsername_multi",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryDuration := time.Minute
			if ctxDeadline, ok := CTX.Deadline(); ok {
				retryDuration = time.Until(ctxDeadline)
			}
			require.EventuallyWithT(t, func(ttt *assert.CollectT) {
				got, listErr := Client.ListIDPLinks(tt.args.ctx, tt.args.req)
				assertErr := assert.NoError
				if tt.wantErr {
					assertErr = assert.Error
				}
				assertErr(ttt, listErr)
				if listErr != nil {
					return
				}
				// always first check length, otherwise its failed anyway
				assert.Len(ttt, got.Result, len(tt.want.Result))
				for i := range tt.want.Result {
					assert.Contains(ttt, got.Result, tt.want.Result[i])
				}
				integration.AssertListDetails(t, tt.want, got)
			}, retryDuration, time.Millisecond*100, "timeout waiting for expected idplinks result")
		})
	}
}

func TestServer_RemoveIDPLink(t *testing.T) {
	orgResp := Instance.CreateOrganization(IamCTX, fmt.Sprintf("ListIDPLinks%d", time.Now().UnixNano()), fmt.Sprintf("%d@mouse.com", time.Now().UnixNano()))

	instanceIdpResp := Instance.AddGenericOAuthProvider(IamCTX, Instance.DefaultOrg.Id)
	userInstanceResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))
	_, err := Instance.CreateUserIDPlink(IamCTX, userInstanceResp.GetUserId(), "external_instance", instanceIdpResp.Id, "externalUsername_instance")
	require.NoError(t, err)
	orgIdpResp := Instance.AddOrgGenericOAuthProvider(IamCTX, orgResp.OrganizationId)
	userOrgResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))
	_, err = Instance.CreateUserIDPlink(IamCTX, userOrgResp.GetUserId(), "external_org", orgIdpResp.Id, "externalUsername_org")
	require.NoError(t, err)

	userNoLinkResp := Instance.CreateHumanUserVerified(IamCTX, orgResp.OrganizationId, fmt.Sprintf("%d@listidplinks.com", time.Now().UnixNano()))

	type args struct {
		ctx context.Context
		req *user.RemoveIDPLinkRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.RemoveIDPLinkResponse
		wantErr bool
	}{
		{
			name: "remove link, no permission",
			args: args{
				UserCTX,
				&user.RemoveIDPLinkRequest{
					UserId:       userOrgResp.GetUserId(),
					IdpId:        orgIdpResp.Id,
					LinkedUserId: "external_org",
				},
			},
			wantErr: true,
		},
		{
			name: "remove link, no permission, org",
			args: args{
				CTX,
				&user.RemoveIDPLinkRequest{
					UserId:       userOrgResp.GetUserId(),
					IdpId:        orgIdpResp.Id,
					LinkedUserId: "external_org",
				},
			},
			wantErr: true,
		},
		{
			name: "remove link, org, ok",
			args: args{
				IamCTX,
				&user.RemoveIDPLinkRequest{
					UserId:       userOrgResp.GetUserId(),
					IdpId:        orgIdpResp.Id,
					LinkedUserId: "external_org",
				},
			},
			want: &user.RemoveIDPLinkResponse{
				Details: &object.Details{
					ResourceOwner: orgResp.GetOrganizationId(),
					ChangeDate:    timestamppb.Now(),
				},
			},
		},
		{
			name: "remove link, instance, ok",
			args: args{
				IamCTX,
				&user.RemoveIDPLinkRequest{
					UserId:       userInstanceResp.GetUserId(),
					IdpId:        instanceIdpResp.Id,
					LinkedUserId: "external_instance",
				},
			},
			want: &user.RemoveIDPLinkResponse{
				Details: &object.Details{
					ResourceOwner: orgResp.GetOrganizationId(),
					ChangeDate:    timestamppb.Now(),
				},
			},
		},
		{
			name: "remove link, no link, error",
			args: args{
				IamCTX,
				&user.RemoveIDPLinkRequest{
					UserId: userNoLinkResp.GetUserId(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.RemoveIDPLink(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			integration.AssertDetails(t, tt.want, got)
		})
	}
}
