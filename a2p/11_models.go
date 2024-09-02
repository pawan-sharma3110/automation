package a2p

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CustomerProfileData struct {
	SID            string `json:"customer_profile_sid"`
	FriendlyName   string `json:"friendly_name"`
	Email          string `json:"email"`
	StatusCallback string `json:"status_callback"`
	PolicySid      string `json:"policy_sid"`
}

type BusinessInfoData struct {
	SID                        string `json:"end_user_sid"`
	BusinessName               string `json:"business_name"`
	SocialMediaProfileUrls     string `json:"social_media_profile_urls"`
	WebsiteUrl                 string `json:"website_url"`
	BusinessRegionsOfOperation string `json:"business_regions_of_operation"`
	BusinessType               string `json:"business_type"`
	BusinessRegistrationId     string `json:"business_registration_identifier"`
	BusinessIdentity           string `json:"business_identity"`
	BusinessIndustry           string `json:"business_industry"`
	BusinessRegistrationNumber string `json:"business_registration_number"`
}

type EndUserAuthorizedRep1BusinessInfoData struct {
	SID           string `json:"end_user_rep1_sid"`
	Type          string `json:"type"`
	Position      string `json:"position"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone"`
	BusinessTitle string `json:"business_name"`
	FriendlyName  string `json:"friendly_name"`
}

type EndUserAssignmentData struct {
	SID                string `json:"assignment_resource_sid"`
	CustomerProfileSid string `json:"customer_profile_sid"`
	EndUserSid         string `json:"end_user_sid"`
}

type AddressData struct {
	SID             string `json:"address_sid"`
	PathAccountSid  string `json:"path_account_sid"`
	CustomerName    string `json:"customer_name"`
	Street          string `json:"street"`
	City            string `json:"city"`
	Region          string `json:"region"`
	PostalCode      string `json:"postal_code"`
	IsoCountry      string `json:"iso_country"`
	FriendlyName    string `json:"friendly_name"`
	StreetSecondary string `json:"street_secondary"`
}

type SupportingDocumentData struct {
	SID          string `json:"supporting_document_sid"`
	FriendlyName string `json:"friendly_name"`
	AddressSid   string `json:"address_sid"`
}

type TrustProductData struct {
	SID            string `json:"trust_product_sid"`
	FriendlyName   string `json:"friendly_name"`
	Email          string `json:"email"`
	PolicySid      string `json:"policy_sid"`
	StatusCallback string `json:"status_callback"`
}

type EndUserMessagingProfileData struct {
	SID           string `json:"end_user_messaging_profile_sid"`
	CompanyType   string `json:"company_type"`
	StockExchange string `json:"stock_exchange"`
	StockTicker   string `json:"stock_ticker"`
}

/*
Note :
The customer_profile_bundle_sid is the SID of your customer's Secondary Customer Profile.
The a2p_profile_bundle_sid is the SID of the TrustProduct created SID.
Skip_automatic_sec_vet is an optional Boolean.
*/
type BrandRegistrationData struct {
	CustomerProfileBundleSid string `json:"customer_profile_bundle_sid"`
	A2PProfileBundleSid      string `json:"a2p_profile_bundle_sid"`
}

type MessagingServiceData struct {
	SID               string `json:"messaging_service_sid"`
	FriendlyName      string `json:"friendly_name"`
	InboundRequestUrl string `json:"inbound_request_url"`
	FallbackUrl       string `json:"fallback_url"`
}

type CampaignData struct {
	SID                  string   `json:"campaign_sid"`
	Description          string   `json:"description"`
	Usecase              string   `json:"usecase"`
	CampaignStatus       string   `json:"campaign_status"`
	HasEmbeddedLinks     bool     `json:"has_embedded_links"`
	HasEmbeddedPhone     bool     `json:"has_embedded_phone"`
	MessageSamples       []string `json:"message_samples"`
	MessageFlow          string   `json:"message_flow"`
	BrandRegistrationSid string   `json:"brand_registration_sid"`
}

// Optional : uncomment for step 4.2
type MessagingServiceAdditional struct {
	SID                   string `json:"messaging_service_sid"`
	FriendlyName          string `json:"friendly_name"`
	InboundRequestUrl     string `json:"inbound_request_url"`
	FallbackUrl           string `json:"fallback_url"`
	StatusCallback        string `json:"status_callback"`
	StickySender          bool   `json:"sticky_sender"`
	SmartEncoding         bool   `json:"smart_encoding"`
	MmsConverter          bool   `json:"mms_converter"`
	FallbackToLongCode    bool   `json:"fallback_to_long_code"`
	ScanMessageContent    string `json:"scan_message_content"`
	AreaCodeGeomatch      bool   `json:"area_code_geomatch"`
	ValidityPeriod        int    `json:"validity_period"`
	SynchronousValidation bool   `json:"synchronous_validation"`
	Usecase               string `json:"usecase"`
}

type MessageStatusData struct {
	MessageSid string `json:"message_sid"`
}

type MessageErrorData struct {
	MessageSid string `json:"message_sid"`
}

type AssociatePhoneNumberData struct {
	PhoneNumberSid      string `json:"phone_number_sid"`
	MessagingServiceSid string `json:"messaging_service_sid"`
}

type FinalizeMessagingServiceConfigData struct {
	MessagingServiceSid   string `json:"messaging_service_sid"`
	StatusCallback        string `json:"status_callback"`
	StickySender          bool   `json:"sticky_sender"`
	SmartEncoding         bool   `json:"smart_encoding"`
	MmsConverter          bool   `json:"mms_converter"`
	FallbackToLongCode    bool   `json:"fallback_to_long_code"`
	ScanMessageContent    string `json:"scan_message_content"`
	AreaCodeGeomatch      bool   `json:"area_code_geomatch"`
	ValidityPeriod        int    `json:"validity_period"`
	SynchronousValidation bool   `json:"synchronous_validation"`
}

type FullA2POnboardingParams struct {
	LocationID                    string `json:"location_id"`
	SubaccountID                  string `json:"subaccount_id"`
	TwilioRootUsername            string `json:"root_username"`
	TwilioRootPassword            string `json:"root_password"`
	TwilioUsername                string `json:"subaccount_username"`
	TwilioPassword                string `json:"subaccount_password"`
	FirstName                     string `json:"first_name"`
	LastName                      string `json:"last_name"`
	Website                       string `json:"website"`
	CustomerName                  string `json:"customer_name"`
	Email                         string `json:"customer_email"`
	PhoneNumber                   string `json:"customer_phone_number"`
	Street                        string `json:"street"`
	City                          string `json:"city"`
	Region                        string `json:"region"`
	PostalCode                    string `json:"postal_code"`
	IsoCountry                    string `json:"iso_country"`
	SocialMediaProfileURLs        string `json:"social_media_profile_urls"`
	WebsiteURL                    string `json:"website_url"`
	FriendlyName                  string `json:"friendly_name"`
	BusinessName                  string `json:"business_name"`
	BusinessIndustry              string `json:"business_industry"`
	BusinessType                  string `json:"business_type"`
	BusinessRegistrationId        string `json:"business_registration_identifier"`
	BusinessIdentity              string `json:"business_identity"`
	BusinessRegistrationNumber    string `json:"business_registration_number"`
	RegionOfOperation             string `json:"region_of_operation"`
	TwilioPurchasedPhoneNumber    string `json:"twilio_purchased_phone_number"`
	TwilioPurchasedPhoneNumberSID string `json:"twilio_purchased_phone_number_sid"`
	AuthorizedRepresentativeName  string `json:"authorized_representative_name"`
	AuthorizedRepresentativeTitle string `json:"authorized_representative_title"`
	AuthorizedRepresentativeEmail string `json:"authorized_representative_email"`
	AuthorizedRepresentativePhone string `json:"authorized_representative_phone"`
	EndUserRepOnePosition         string `json:"end_user_rep_one_position"`
	EndUserRepOneFirstName        string `json:"end_user_rep_one_first_name"`
	EndUserRepOneLastName         string `json:"end_user_rep_one_last_name"`
	EndUserRepOneEmail            string `json:"end_user_rep_one_email"`
	EndUserRepOnePhoneNumber      string `json:"end_user_rep_one_phone"`
	EndUserRepOneBusinessTitle    string `json:"end_user_rep_one_business_name"`
	EndUserRepOneFriendlyName     string `json:"end_user_rep_one_friendly_name"`
	UseCase                       string `json:"use_case"`
	AreaCode                      string `json:"area_code"`
	BrandRegistrationSID          string `json:"brand_registration_sid"`
	MessagingServiceSID           string `json:"messaging_service_sid"`
}

func (f *FullA2POnboardingParams) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.CustomerName, validation.Required),
		validation.Field(&f.Email, validation.Required),
		validation.Field(&f.PhoneNumber, validation.Required),
		validation.Field(&f.Street, validation.Required),
		validation.Field(&f.City, validation.Required),
		validation.Field(&f.Region, validation.Required),
		validation.Field(&f.PostalCode, validation.Required),
		validation.Field(&f.IsoCountry, validation.Required),
		validation.Field(&f.SocialMediaProfileURLs, validation.Required),
		validation.Field(&f.WebsiteURL, validation.Required),
		validation.Field(&f.BusinessName, validation.Required),
		validation.Field(&f.BusinessIndustry, validation.Required),
		validation.Field(&f.BusinessType, validation.Required),
		validation.Field(&f.BusinessRegistrationId, validation.Required),
		validation.Field(&f.BusinessIdentity, validation.Required),
		validation.Field(&f.BusinessRegistrationNumber, validation.Required),
		validation.Field(&f.RegionOfOperation, validation.Required),
		validation.Field(&f.TwilioPurchasedPhoneNumber, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeName, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeTitle, validation.Required),
		validation.Field(&f.AuthorizedRepresentativeEmail, validation.Required),
		validation.Field(&f.AuthorizedRepresentativePhone, validation.Required),
		validation.Field(&f.FriendlyName, validation.Required),
		validation.Field(&f.UseCase, validation.Required),
		validation.Field(&f.AreaCode, validation.Required),
	)

}

type FullA2POnboardingResponse struct {
	Message string `json:"message"`
	Data    *A2POnboardingResponse
}

type A2POnboardingResponse struct {
	ID                           string    `json:"id"`
	LocationID                   string    `json:"location_id"`
	SubaccountID                 string    `json:"subaccount_id"`
	TwilioUsername               string    `json:"subaccount_username"`
	TwilioPassword               string    `json:"subaccount_password"`
	TwilioPhoneNumber            string    `json:"twilio_phone_number"`
	TwilioPhoneNumberSID         string    `json:"twilio_phone_number_sid"`
	BrandRegistrationSID         string    `json:"brand_registration_sid"`
	MessagingServiceSID          string    `json:"messaging_service_sid"`
	A2pMessageCampaignSID        string    `json:"a2p_message_campaign_sid"`
	AppliedForBrandRegistration  bool      `json:"applied_for_brand_registration"`
	AppliedForMessagingService   bool      `json:"applied_for_messaging_service"`
	AppliedForA2pMessageCampaign bool      `json:"applied_for_a2p_message_campaign"`
	BrandRegistrationStatus      string    `json:"brand_registration_status"`
	MessagingServiceStatus       string    `json:"messaging_service_status"`
	A2pMessageCampaignStatus     string    `json:"a2p_message_campaign_status"`
	Created_At                   time.Time `json:"created_at"`
	Updated_At                   time.Time `json:"updated_at"`
}
