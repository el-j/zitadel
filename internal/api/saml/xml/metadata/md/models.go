// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for urn:oasis:names:tc:SAML:2.0:metadata
package md

import (
	"encoding/xml"
	"github.com/caos/zitadel/internal/api/saml/xml/metadata/saml"
	"github.com/caos/zitadel/internal/api/saml/xml/metadata/xml_dsig"
)

// Element
type Extensions struct {
	XMLName xml.Name `xml:"Extensions"`
}

// Element
type EntitiesDescriptor struct {
	XMLName xml.Name `xml:"EntitiesDescriptor"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	EntityDescriptor []EntityDescriptorType `xml:"EntityDescriptor"`

	EntitiesDescriptor []EntitiesDescriptorType `xml:"EntitiesDescriptor"`
}

// Element
type EntityDescriptor struct {
	XMLName                      xml.Name                          `xml:"EntityDescriptor"`
	EntityID                     EntityIDType                      `xml:"entityID,attr"`
	ValidUntil                   string                            `xml:"validUntil,attr,omitempty"`
	CacheDuration                string                            `xml:"cacheDuration,attr,omitempty"`
	Id                           string                            `xml:"ID,attr,omitempty"`
	Signature                    *xml_dsig.SignatureType           `xml:"Signature"`
	Extensions                   *ExtensionsType                   `xml:"Extensions"`
	Organization                 *OrganizationType                 `xml:"Organization"`
	ContactPerson                []ContactType                     `xml:"ContactPerson"`
	AdditionalMetadataLocation   []AdditionalMetadataLocationType  `xml:"AdditionalMetadataLocation"`
	RoleDescriptor               *RoleDescriptorType               `xml:"RoleDescriptor,omitempty"`
	IDPSSODescriptor             *IDPSSODescriptorType             `xml:"IDPSSODescriptor,omitempty"`
	SPSSODescriptor              *SPSSODescriptorType              `xml:"SPSSODescriptor,omitempty"`
	AuthnAuthorityDescriptor     *AuthnAuthorityDescriptorType     `xml:"AuthnAuthorityDescriptor,omitempty"`
	AttributeAuthorityDescriptor *AttributeAuthorityDescriptorType `xml:"AttributeAuthorityDescriptor,omitempty"`
	PDPDescriptor                *PDPDescriptorType                `xml:"PDPDescriptor,omitempty"`
	AffiliationDescriptor        *AffiliationDescriptorType        `xml:"AffiliationDescriptor,omitempty"`
}

// Element
type Organization struct {
	XMLName xml.Name `xml:"Organization"`

	Extensions *ExtensionsType `xml:"Extensions"`

	OrganizationName []LocalizedNameType `xml:"OrganizationName"`

	OrganizationDisplayName []LocalizedNameType `xml:"OrganizationDisplayName"`

	OrganizationURL []LocalizedURIType `xml:"OrganizationURL"`
}

// Element
type OrganizationName struct {
	XMLName xml.Name `xml:"OrganizationName"`

	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}

// Element
type OrganizationDisplayName struct {
	XMLName xml.Name `xml:"OrganizationDisplayName"`

	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}

// Element
type OrganizationURL struct {
	XMLName xml.Name `xml:"OrganizationURL"`

	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}

// Element
type ContactPerson struct {
	XMLName xml.Name `xml:"ContactPerson"`

	ContactType ContactTypeType `xml:"contactType,attr"`

	Extensions *ExtensionsType `xml:"Extensions"`

	Company string `xml:"Company"`

	GivenName string `xml:"GivenName"`

	SurName string `xml:"SurName"`

	EmailAddress []string `xml:"EmailAddress"`

	TelephoneNumber []string `xml:"TelephoneNumber"`
}

// Element
type Company struct {
	XMLName xml.Name `xml:"Company"`

	Text string `xml:",chardata"`
}

// Element
type GivenName struct {
	XMLName xml.Name `xml:"GivenName"`

	Text string `xml:",chardata"`
}

// Element
type SurName struct {
	XMLName xml.Name `xml:"SurName"`

	Text string `xml:",chardata"`
}

// Element
type EmailAddress struct {
	XMLName xml.Name `xml:"EmailAddress"`

	Text string `xml:",chardata"`
}

// Element
type TelephoneNumber struct {
	XMLName xml.Name `xml:"TelephoneNumber"`

	Text string `xml:",chardata"`
}

// Element
type AdditionalMetadataLocation struct {
	XMLName xml.Name `xml:"AdditionalMetadataLocation"`

	Namespace string `xml:"namespace,attr"`

	Text string `xml:",chardata"`
}

// Element
type RoleDescriptor struct {
	XMLName xml.Name `xml:"RoleDescriptor"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type KeyDescriptor struct {
	XMLName xml.Name `xml:"KeyDescriptor"`

	Use KeyTypes `xml:"use,attr,omitempty"`

	KeyInfo xml_dsig.KeyInfoType `xml:"KeyInfo"`

	EncryptionMethod []EncryptionMethodType `xml:"EncryptionMethod"`
}

// Element
type EncryptionMethod struct {
	XMLName xml.Name `xml:"EncryptionMethod"`

	Algorithm string `xml:"Algorithm,attr"`

	KeySize *KeySizeType `xml:"KeySize"`

	OAEPparams string `xml:"OAEPparams"`
}

// Element
type ArtifactResolutionService struct {
	XMLName xml.Name `xml:"ArtifactResolutionService"`

	Index string `xml:"index,attr"`

	IsDefault string `xml:"isDefault,attr,omitempty"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type SingleLogoutService struct {
	XMLName xml.Name `xml:"SingleLogoutService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type ManageNameIDService struct {
	XMLName xml.Name `xml:"ManageNameIDService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type NameIDFormat struct {
	XMLName xml.Name `xml:"NameIDFormat"`

	Text string `xml:",chardata"`
}

// Element
type IDPSSODescriptor struct {
	XMLName xml.Name `xml:"IDPSSODescriptor"`

	WantAuthnRequestsSigned string `xml:"WantAuthnRequestsSigned,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	SingleSignOnService []EndpointType `xml:"SingleSignOnService"`

	NameIDMappingService []EndpointType `xml:"NameIDMappingService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	AttributeProfile []string `xml:"AttributeProfile"`

	Attribute []saml.AttributeType `xml:"Attribute"`

	ArtifactResolutionService []IndexedEndpointType `xml:"ArtifactResolutionService"`

	SingleLogoutService []EndpointType `xml:"SingleLogoutService"`

	ManageNameIDService []EndpointType `xml:"ManageNameIDService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type SingleSignOnService struct {
	XMLName xml.Name `xml:"SingleSignOnService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type NameIDMappingService struct {
	XMLName xml.Name `xml:"NameIDMappingService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type AssertionIDRequestService struct {
	XMLName xml.Name `xml:"AssertionIDRequestService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type AttributeProfile struct {
	XMLName xml.Name `xml:"AttributeProfile"`

	Text string `xml:",chardata"`
}

// Element
type SPSSODescriptor struct {
	XMLName xml.Name `xml:"SPSSODescriptor"`

	AuthnRequestsSigned string `xml:"AuthnRequestsSigned,attr,omitempty"`

	WantAssertionsSigned string `xml:"WantAssertionsSigned,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AssertionConsumerService []IndexedEndpointType `xml:"AssertionConsumerService"`

	AttributeConsumingService []AttributeConsumingServiceType `xml:"AttributeConsumingService"`

	ArtifactResolutionService []IndexedEndpointType `xml:"ArtifactResolutionService"`

	SingleLogoutService []EndpointType `xml:"SingleLogoutService"`

	ManageNameIDService []EndpointType `xml:"ManageNameIDService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type AssertionConsumerService struct {
	XMLName xml.Name `xml:"AssertionConsumerService"`

	Index string `xml:"index,attr"`

	IsDefault string `xml:"isDefault,attr,omitempty"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type AttributeConsumingService struct {
	XMLName xml.Name `xml:"AttributeConsumingService"`

	Index uint64 `xml:"index,attr"`

	IsDefault bool `xml:"isDefault,attr,omitempty"`

	ServiceName []LocalizedNameType `xml:"ServiceName"`

	ServiceDescription []LocalizedNameType `xml:"ServiceDescription"`

	RequestedAttribute []RequestedAttributeType `xml:"RequestedAttribute"`
}

// Element
type ServiceName struct {
	XMLName xml.Name `xml:"ServiceName"`

	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}

// Element
type ServiceDescription struct {
	XMLName xml.Name `xml:"ServiceDescription"`

	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}

// Element
type RequestedAttribute struct {
	XMLName xml.Name `xml:"RequestedAttribute"`

	IsRequired string `xml:"isRequired,attr,omitempty"`

	Name string `xml:"Name,attr"`

	NameFormat string `xml:"NameFormat,attr,omitempty"`

	FriendlyName string `xml:"FriendlyName,attr,omitempty"`

	AttributeValue []saml.string `xml:",any"`
}

// Element
type AuthnAuthorityDescriptor struct {
	XMLName xml.Name `xml:"AuthnAuthorityDescriptor"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AuthnQueryService []EndpointType `xml:"AuthnQueryService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type AuthnQueryService struct {
	XMLName xml.Name `xml:"AuthnQueryService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type PDPDescriptor struct {
	XMLName xml.Name `xml:"PDPDescriptor"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AuthzService []EndpointType `xml:"AuthzService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type AuthzService struct {
	XMLName xml.Name `xml:"AuthzService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type AttributeAuthorityDescriptor struct {
	XMLName xml.Name `xml:"AttributeAuthorityDescriptor"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AttributeService []EndpointType `xml:"AttributeService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	AttributeProfile []string `xml:"AttributeProfile"`

	Attribute []saml.AttributeType `xml:"Attribute"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`
}

// Element
type AttributeService struct {
	XMLName xml.Name `xml:"AttributeService"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
}

// Element
type AffiliationDescriptor struct {
	XMLName xml.Name `xml:"AffiliationDescriptor"`

	AffiliationOwnerID EntityIDType `xml:"affiliationOwnerID,attr"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	AffiliateMember []EntityIDType `xml:"AffiliateMember"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`
}

// Element
type AffiliateMember struct {
	XMLName xml.Name `xml:"AffiliateMember"`

	Text string `xml:",chardata"`
}

// XSD ComplexType declarations

type LocalizedNameType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type LocalizedURIType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ExtensionsType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type EndpointType struct {
	XMLName xml.Name

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type IndexedEndpointType struct {
	XMLName xml.Name

	Index string `xml:"index,attr"`

	IsDefault string `xml:"isDefault,attr,omitempty"`

	Binding string `xml:"Binding,attr"`

	Location string `xml:"Location,attr"`

	ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntitiesDescriptorType struct {
	XMLName xml.Name

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	EntityDescriptor []EntityDescriptorType `xml:"EntityDescriptor"`

	EntitiesDescriptor []EntitiesDescriptorType `xml:"EntitiesDescriptor"`

	InnerXml string `xml:",innerxml"`
}

type EntityDescriptorType struct {
	XMLName xml.Name

	EntityID EntityIDType `xml:"entityID,attr"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	AdditionalMetadataLocation []AdditionalMetadataLocationType `xml:"AdditionalMetadataLocation"`

	AffiliationDescriptor *AffiliationDescriptorType `xml:"AffiliationDescriptor"`

	InnerXml string `xml:",innerxml"`
}

type OrganizationType struct {
	XMLName xml.Name

	Extensions *ExtensionsType `xml:"Extensions"`

	OrganizationName []LocalizedNameType `xml:"OrganizationName"`

	OrganizationDisplayName []LocalizedNameType `xml:"OrganizationDisplayName"`

	OrganizationURL []LocalizedURIType `xml:"OrganizationURL"`

	InnerXml string `xml:",innerxml"`
}

type ContactType struct {
	XMLName xml.Name

	ContactType ContactTypeType `xml:"contactType,attr"`

	Extensions *ExtensionsType `xml:"Extensions"`

	Company string `xml:"Company"`

	GivenName string `xml:"GivenName"`

	SurName string `xml:"SurName"`

	EmailAddress []string `xml:"EmailAddress"`

	TelephoneNumber []string `xml:"TelephoneNumber"`

	InnerXml string `xml:",innerxml"`
}

type AdditionalMetadataLocationType struct {
	XMLName xml.Name

	Namespace string `xml:"namespace,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type RoleDescriptorType struct {
	XMLName xml.Name

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type KeyDescriptorType struct {
	XMLName xml.Name

	Use KeyTypes `xml:"use,attr,omitempty"`

	KeyInfo xml_dsig.KeyInfoType `xml:"KeyInfo"`

	EncryptionMethod []EncryptionMethodType `xml:"EncryptionMethod"`

	InnerXml string `xml:",innerxml"`
}

type SSODescriptorType struct {
	XMLName xml.Name

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	ArtifactResolutionService []IndexedEndpointType `xml:"ArtifactResolutionService"`

	SingleLogoutService []EndpointType `xml:"SingleLogoutService"`

	ManageNameIDService []EndpointType `xml:"ManageNameIDService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type IDPSSODescriptorType struct {
	XMLName xml.Name

	WantAuthnRequestsSigned string `xml:"WantAuthnRequestsSigned,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	SingleSignOnService []EndpointType `xml:"SingleSignOnService"`

	NameIDMappingService []EndpointType `xml:"NameIDMappingService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	AttributeProfile []string `xml:"AttributeProfile"`

	Attribute []saml.AttributeType `xml:"Attribute"`

	ArtifactResolutionService []IndexedEndpointType `xml:"ArtifactResolutionService"`

	SingleLogoutService []EndpointType `xml:"SingleLogoutService"`

	ManageNameIDService []EndpointType `xml:"ManageNameIDService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type SPSSODescriptorType struct {
	XMLName xml.Name

	AuthnRequestsSigned string `xml:"AuthnRequestsSigned,attr,omitempty"`

	WantAssertionsSigned string `xml:"WantAssertionsSigned,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AssertionConsumerService []IndexedEndpointType `xml:"AssertionConsumerService"`

	AttributeConsumingService []AttributeConsumingServiceType `xml:"AttributeConsumingService"`

	ArtifactResolutionService []IndexedEndpointType `xml:"ArtifactResolutionService"`

	SingleLogoutService []EndpointType `xml:"SingleLogoutService"`

	ManageNameIDService []EndpointType `xml:"ManageNameIDService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type AttributeConsumingServiceType struct {
	XMLName xml.Name

	Index uint64 `xml:"index,attr"`

	IsDefault bool `xml:"isDefault,attr,omitempty"`

	ServiceName []LocalizedNameType `xml:"ServiceName"`

	ServiceDescription []LocalizedNameType `xml:"ServiceDescription"`

	RequestedAttribute []RequestedAttributeType `xml:"RequestedAttribute"`

	InnerXml string `xml:",innerxml"`
}

type RequestedAttributeType struct {
	XMLName xml.Name

	IsRequired string `xml:"isRequired,attr,omitempty"`

	Name string `xml:"Name,attr"`

	NameFormat string `xml:"NameFormat,attr,omitempty"`

	FriendlyName string `xml:"FriendlyName,attr,omitempty"`

	AttributeValue []saml.string `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type AuthnAuthorityDescriptorType struct {
	XMLName xml.Name

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AuthnQueryService []EndpointType `xml:"AuthnQueryService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type PDPDescriptorType struct {
	XMLName xml.Name

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AuthzService []EndpointType `xml:"AuthzService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type AttributeAuthorityDescriptorType struct {
	XMLName xml.Name

	Id string `xml:"ID,attr,omitempty"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	ProtocolSupportEnumeration AnyURIListType `xml:"protocolSupportEnumeration,attr"`

	ErrorURL string `xml:"errorURL,attr,omitempty"`

	AttributeService []EndpointType `xml:"AttributeService"`

	AssertionIDRequestService []EndpointType `xml:"AssertionIDRequestService"`

	NameIDFormat []string `xml:"NameIDFormat"`

	AttributeProfile []string `xml:"AttributeProfile"`

	Attribute []saml.AttributeType `xml:"Attribute"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	Organization *OrganizationType `xml:"Organization"`

	ContactPerson []ContactType `xml:"ContactPerson"`

	InnerXml string `xml:",innerxml"`
}

type AffiliationDescriptorType struct {
	XMLName xml.Name

	AffiliationOwnerID EntityIDType `xml:"affiliationOwnerID,attr"`

	ValidUntil string `xml:"validUntil,attr,omitempty"`

	CacheDuration string `xml:"cacheDuration,attr,omitempty"`

	Id string `xml:"ID,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Extensions *ExtensionsType `xml:"Extensions"`

	AffiliateMember []EntityIDType `xml:"AffiliateMember"`

	KeyDescriptor []KeyDescriptorType `xml:"KeyDescriptor"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type EntityIDType string

type ContactTypeType string

const ContactTypeTypeTechnical ContactTypeType = "technical"

const ContactTypeTypeSupport ContactTypeType = "support"

const ContactTypeTypeAdministrative ContactTypeType = "administrative"

const ContactTypeTypeBilling ContactTypeType = "billing"

const ContactTypeTypeOther ContactTypeType = "other"

type AnyURIListType string

type KeyTypes string

const KeyTypesEncryption KeyTypes = "encryption"

const KeyTypesSigning KeyTypes = "signing"
