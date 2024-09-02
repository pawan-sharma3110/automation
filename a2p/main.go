package a2p

import (
	// "crm/internal/fmt"Printlnrors"
	"errors"
	"fmt"
	"time"

	"github.com/twilio/twilio-go"
	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

type A2PService struct {
	client *twilio.RestClient
}

func NewA2PServiceInstance(sid, token string) *A2PService {
	return &A2PService{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: sid,
			Password: token,
		}),
	}
}

var (
	ErrCreateSubaccount               = errors.New("create a subaccount first before proceeding")
	ErrPurchasePhoneNumber            = errors.New("purchase a Twilio phone number first before proceeding")
	ErrGetPhoneNumberSID              = errors.New("get the Twilio phone number SID first before proceeding")
	ErrGetTwilioUsername              = errors.New("get the Twilio root username first before proceeding")
	ErrGetTwilioPassword              = errors.New("get the Twilio root password first before proceeding")
	ErrBrandRegistrationCheckTimedOut = errors.New("checking brand registration timed out after 48 hours")
)

func (s *A2PService) OnboardCustomer(params *FullA2POnboardingParams) (FullA2POnboardingResponse, error) {

	if params.SubaccountID == "" {
		return FullA2POnboardingResponse{}, ErrCreateSubaccount
	}

	if params.TwilioPurchasedPhoneNumber == "" {
		return FullA2POnboardingResponse{}, ErrPurchasePhoneNumber
	}

	if params.TwilioPurchasedPhoneNumberSID == "" {
		return FullA2POnboardingResponse{}, ErrGetPhoneNumberSID
	}

	if params.TwilioUsername == "" {
		return FullA2POnboardingResponse{}, ErrGetTwilioUsername
	}

	if params.TwilioPassword == "" {
		return FullA2POnboardingResponse{}, ErrGetTwilioPassword
	}

	fmt.Println("Starting onboarding process for", params.FriendlyName)

	// Stage 2.1: Create a secondary customer profile
	customerProfileSid, err := s.CreateSecondaryCustomerProfile(CustomerProfileData{
		FriendlyName:   params.FriendlyName,
		Email:          params.Email,
		PolicySid:      "RNdfbf3fae0e1107f8aded0e7cead80bf5",
		StatusCallback: "www.demo.com/callback/status",
	})

	if err != nil {
		fmt.Println("CreateSecondaryCustomerProfile", "error at stage 2.1", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.2: Create an EndUser Business Information resource
	endUserBusinessInfoSID, err := s.CreateEndUserBusinessInfo(BusinessInfoData{
		BusinessName:               params.BusinessName,
		SocialMediaProfileUrls:     params.SocialMediaProfileURLs,
		WebsiteUrl:                 params.WebsiteURL,
		BusinessRegionsOfOperation: params.RegionOfOperation,
		BusinessType:               params.BusinessType,
		BusinessRegistrationId:     params.BusinessRegistrationId,
		BusinessIdentity:           params.BusinessIdentity,
		BusinessIndustry:           params.BusinessIndustry,
		BusinessRegistrationNumber: params.BusinessRegistrationNumber,
	})
	if err != nil {
		fmt.Println("CreateEndUserBusinessInfo", "error at stage 2.2", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.3: Attach EndUser to the Secondary Customer Profile
	// attachEndUserToProfileSID
	_, err = s.AttachEndUserToProfile(EndUserAssignmentData{
		CustomerProfileSid: customerProfileSid,
		EndUserSid:         endUserBusinessInfoSID,
	})
	if err != nil {
		return FullA2POnboardingResponse{}, fmt.Errorf("error at stage 2.3: %w", err)
	}

	// Stage 2.4. Create an EndUser resource of type: authorized_representative_1
	endUserAuthorizedRep1SID, err := s.CreateEndUserAuthorizedRep1(EndUserAuthorizedRep1BusinessInfoData{
		Type:          "authorized_representative_1",
		FirstName:     params.EndUserRepOneFirstName,
		LastName:      params.EndUserRepOneLastName,
		Email:         params.EndUserRepOneEmail,
		PhoneNumber:   params.EndUserRepOneEmail,
		Position:      params.EndUserRepOnePosition,
		BusinessTitle: params.EndUserRepOneBusinessTitle,
		FriendlyName:  fmt.Sprintf("%s - Authorized Representative 1", params.CustomerName),
	})
	if err != nil {
		fmt.Println("CreateEndUserAuthorizedRep1", "error at stage 2.4", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.5: Attach EndUser to the Secondary Customer Profile
	// attachEndUserToProfileSID
	_, err = s.AttachEndUserAuthorizedRep1ToProfile(EndUserAssignmentData{
		CustomerProfileSid: customerProfileSid,
		EndUserSid:         endUserAuthorizedRep1SID,
	})
	if err != nil {
		fmt.Println("AttachEndUserAuthorizedRep1ToProfile", "error at stage 2.5", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.6 Create An Address Resource and returns address sid
	addressSID, err := s.CreateAddressResource(AddressData{
		PathAccountSid: params.TwilioUsername,
		CustomerName:   params.CustomerName,
		Street:         params.Street,
		City:           params.City,
		Region:         params.Region,
		PostalCode:     params.PostalCode,
		IsoCountry:     params.IsoCountry,
		FriendlyName:   fmt.Sprintf("%s - Address Resource", params.CustomerName),
	})
	if err != nil {
		fmt.Println("CreateAddressResource", "error at stage 2.6", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.7 Create a supporting document resource and returns supporting_document_sid
	supportingDocumentSID, err := s.CreateSupportingDocumentResource(SupportingDocumentData{
		FriendlyName: fmt.Sprintf("%s - Business License Document", params.CustomerName),
		AddressSid:   addressSID,
	})
	if err != nil {
		fmt.Println("CreateSupportingDocument", "error at stage 2.7", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.8 Attach the supporting document to the Secondary Customer Profile
	//attachSupportingDocumentToProfileSID
	_, err = s.AttachSupportingDocumentToProfile(customerProfileSid, &supportingDocumentSID)
	if err != nil {
		fmt.Println("AttachSupportingDocumentToProfile", "error at stage 2.8", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.9. Evaluate the Secondary Customer Profile
	//evaluateSecondaryCustomerProfileSID
	_, err = s.EvaluateSecondaryCustomerProfile(customerProfileSid)
	if err != nil {
		fmt.Println("EvaluateSecondaryCustomerProfile", "error at stage 2.9", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 2.10. Submit the Secondary Customer Profile for review  - status must be set to pending-review
	// submitSecondaryCustomerProfileForReviewSID
	_, err = s.SubmitSecondaryCustomerProfileForReview(customerProfileSid)
	if err != nil {
		fmt.Println("SubmitSecondaryCustomerProfileForReview", "error at stage 2.10", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.1: Create a TrustProduct Resource
	trustProductSID, err := s.CreateTrustProduct(TrustProductData{
		FriendlyName:   params.FriendlyName,
		PolicySid:      "RNdfbf3fae0e1107f8aded0e7cead80bf5",
		Email:          params.Email,
		StatusCallback: "www.demo.com/callback/status",
	})
	if err != nil {
		fmt.Println("CreateTrustProduct", "error at stage 3.1", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.2: Create an EndUser Resource of Type us_a2p_messaging_profile_information
	endUserMessagingProfileSID, err := s.CreateEndUserMessagingProfile(EndUserMessagingProfileData{
		CompanyType:   params.BusinessType,
		StockExchange: "",
		StockTicker:   "",
	})
	if err != nil {
		fmt.Println("CreateEndUserMessagingProfile", "error at stage 3.2", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.3: Attach the EndUser to the TrustProduct
	//attachEndUserToTrustProductSID
	_, err = s.AttachEndUserToTrustProduct(trustProductSID, endUserMessagingProfileSID)
	if err != nil {
		fmt.Println("AttachEndUserToTrustProduct", "error at stage 3.3", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.4: Attach the Secondary Customer Profile to the TrustProduct
	//attachSecondaryCustomerProfileToTrustProductSID
	_, err = s.AttachSecondaryCustomerProfileToTrustProduct(trustProductSID, customerProfileSid)
	if err != nil {
		fmt.Println("AttachSecondaryCustomerProfileToTrustProduct", "error at stage 3.4", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.5: Evaluate the TrustProduct
	// evaluateTrustProductSID
	_, err = s.EvaluateTrustProduct(trustProductSID, "RNdfbf3fae0e1107f8aded0e7cead80bf5")
	if err != nil {
		fmt.Println("EvaluateTrustProduct", "error at stage 3.5", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 3.6: Submit the TrustProduct for Review  - status must be set to pending-review
	// submitTrustProductForReviewSID
	_, err = s.SubmitTrustProductForReview(trustProductSID)
	if err != nil {
		fmt.Println("SubmitTrustProductForReview", "error at stage 3.6", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 4.1: Create a BrandRegistration
	brandRegistrationSID, brandRegistrationStatus, err := s.CreateBrandRegistration(BrandRegistrationData{
		CustomerProfileBundleSid: customerProfileSid,
		A2PProfileBundleSid:      trustProductSID,
	})

	if err != nil {
		fmt.Println("CreateBrandRegistration", "error at stage 4.1", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 5.1: Create a MessagingService Resource - This will return MessageServiceSID
	messagingServiceSID, err := s.CreateMessagingService(MessagingServiceData{
		FriendlyName:      params.FriendlyName,
		InboundRequestUrl: "https://www.example.com/inbound-messages-webhook",
		FallbackUrl:       "https://www.example.com/fallback",
	})

	if err != nil {
		fmt.Println("CreateMessagingService", "error at stage 5.1", err)
		return FullA2POnboardingResponse{}, err
	}

	fmt.Println("|---------------*********************************************************----------------|")
	fmt.Println("Brand Registration SID:", brandRegistrationSID)
	fmt.Println("Brand Registration Current Status:", brandRegistrationStatus)
	fmt.Println("|----------------********************************************************----------------|")

	params.MessagingServiceSID = messagingServiceSID

	return FullA2POnboardingResponse{
		Message: "Brand Registration Created Successfully",
		Data: &A2POnboardingResponse{
			LocationID:                   params.LocationID,
			SubaccountID:                 params.SubaccountID,
			TwilioUsername:               params.TwilioUsername,
			TwilioPassword:               params.TwilioPassword,
			BrandRegistrationSID:         brandRegistrationSID,
			MessagingServiceSID:          messagingServiceSID,
			A2pMessageCampaignSID:        "not submitted",
			BrandRegistrationStatus:      brandRegistrationStatus,
			TwilioPhoneNumber:            params.TwilioPurchasedPhoneNumber,
			TwilioPhoneNumberSID:         params.TwilioPurchasedPhoneNumberSID,
			AppliedForBrandRegistration:  true,
			AppliedForMessagingService:   false,
			AppliedForA2pMessageCampaign: false,
		},
	}, nil

}

func (s *A2PService) CompleteOnboarding(params *FullA2POnboardingParams, brandRegistrationSID string) (FullA2POnboardingResponse, error) {

	twilioPhoneSID, err := s.GetPhoneNumberSID(params.TwilioPurchasedPhoneNumber)
	if err != nil {
		fmt.Println("GetPhoneNumberSID", "error at stage 6.0", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 6.1: Add a Phone Number to the Messaging Service , once you have a phone number, you can associate it with the messaging service
	_, err = s.AddPhoneNumberToMessagingService(params.MessagingServiceSID, &messaging.CreatePhoneNumberParams{
		PhoneNumberSid: &twilioPhoneSID,
	})

	if err != nil {
		fmt.Println("AddPhoneNumberToMessagingService", "error at stage 7.1", err)
		return FullA2POnboardingResponse{}, err
	}

	// Stage 7.1: Create the A2P Campaign
	campaignSID, err := s.CreateA2PCampaign(params.MessagingServiceSID, CampaignData{
		BrandRegistrationSid: brandRegistrationSID,
	})

	if err != nil {
		fmt.Println("CreateA2PCampaign", "error at stage 6.1", err)
		return FullA2POnboardingResponse{}, err
	}

	return FullA2POnboardingResponse{
		Message: "Success !! Proceed To Register A2P Campaign Once Brand Registration is Approved",
		Data: &A2POnboardingResponse{
			LocationID:              params.LocationID,
			SubaccountID:            params.SubaccountID,
			TwilioUsername:          params.TwilioUsername,
			TwilioPassword:          params.TwilioPassword,
			BrandRegistrationSID:    brandRegistrationSID,
			MessagingServiceSID:     params.MessagingServiceSID,
			A2pMessageCampaignSID:   campaignSID,
			BrandRegistrationStatus: "Approved",
		},
	}, nil
}

func (s *A2PService) processRegistrationStatus(status string, params *FullA2POnboardingParams, sid string) (FullA2POnboardingResponse, error) {
	switch status {
	case "APPROVED":
		return s.CompleteOnboarding(params, sid)
	case "FAILED", "IN_REVIEW", "PENDING", "DELETED":
		return FullA2POnboardingResponse{Message: fmt.Sprintf("Brand registration is %s", status)}, nil
	default:
		return FullA2POnboardingResponse{}, fmt.Errorf("unknown status: %s", status)
	}
}

func (s *A2PService) MonitorBrandRegistration(brandRegistrationSID string, params *FullA2POnboardingParams) (FullA2POnboardingResponse, error) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	timeout := time.After(48 * time.Hour)

	for {
		select {
		case <-timeout:
			return FullA2POnboardingResponse{Message: "Brand registration checking timed out"}, ErrBrandRegistrationCheckTimedOut
		case <-ticker.C:
			status, err := s.FetchBrandRegistration(brandRegistrationSID)
			if err != nil {
				fmt.Println("CheckBrandRegistrationStatus", "error", err)
				continue
			}

			// Process based on registration status
			response, err := s.processRegistrationStatus(status, params, brandRegistrationSID)
			if err != nil {
				fmt.Println("ProcessRegistrationStatus", "error", err)
				continue
			}
			return response, nil
		}
	}
}
