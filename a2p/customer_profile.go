package a2p

import (
	"fmt"

	api "github.com/twilio/twilio-go/rest/api/v2010"
	trusthub "github.com/twilio/twilio-go/rest/trusthub/v1"
)

// Step 2.1: Create a Secondary Customer Profile
func (s *A2PService) CreateSecondaryCustomerProfile(data CustomerProfileData) (string, error) {
	params := &trusthub.CreateCustomerProfileParams{}
	params.SetPolicySid(data.PolicySid)
	params.SetFriendlyName(data.FriendlyName)
	params.SetEmail(data.Email)
	if data.StatusCallback != "" {
		params.SetStatusCallback(data.StatusCallback)
	}

	resp, err := s.client.TrusthubV1.CreateCustomerProfile(params)
	if err != nil {
		return "", fmt.Errorf("failed to create customer profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.2: Create an EndUser Resource of Type
func (s *A2PService) CreateEndUserBusinessInfo(data BusinessInfoData) (string, error) {
	params := &trusthub.CreateEndUserParams{}
	params.SetAttributes(map[string]interface{}{
		"business_name":                    data.BusinessName,
		"social_media_profile_urls":        data.SocialMediaProfileUrls,
		"website_url":                      data.WebsiteUrl,
		"business_regions_of_operation":    data.BusinessRegionsOfOperation,
		"business_type":                    data.BusinessType,
		"business_registration_identifier": data.BusinessRegistrationId,
		"business_identity":                data.BusinessIdentity,
		"business_industry":                data.BusinessIndustry,
		"business_registration_number":     data.BusinessRegistrationNumber,
	})
	params.SetFriendlyName(fmt.Sprintf("%s - Business Information EndUser resource", data.BusinessName))
	params.SetType("customer_profile_business_information")

	resp, err := s.client.TrusthubV1.CreateEndUser(params)
	if err != nil {
		return "", fmt.Errorf("failed to create EndUser business information: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.3: Attach the EndUser to the Secondary Customer Profile
func (s *A2PService) AttachEndUserToProfile(data EndUserAssignmentData) (string, error) {
	params := &trusthub.CreateCustomerProfileEntityAssignmentParams{}
	params.SetObjectSid(data.EndUserSid)

	resp, err := s.client.TrusthubV1.CreateCustomerProfileEntityAssignment(data.CustomerProfileSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to attach EndUser to customer profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.4. Create an EndUser resource of type: authorized_representative_1
func (s *A2PService) CreateEndUserAuthorizedRep1(data EndUserAuthorizedRep1BusinessInfoData) (string, error) {
	params := &trusthub.CreateEndUserParams{
		Type:         &data.Type,
		FriendlyName: &data.FriendlyName,
	}
	params.SetAttributes(map[string]interface{}{
		"job_position":   data.Position,
		"last_name":      data.LastName,
		"phone_number":   data.PhoneNumber,
		"first_name":     data.FirstName,
		"email":          data.Email,
		"business_title": data.BusinessTitle,
	})
	params.SetFriendlyName(fmt.Sprintf("%s - Authorized Representative 1", data.BusinessTitle))
	params.SetType("authorized_representative_1")

	resp, err := s.client.TrusthubV1.CreateEndUser(params)
	if err != nil {
		return "", fmt.Errorf("failed to create EndUser authorized representative 1: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.5 Attach the EndUser to the Secondary Customer Profile
func (s *A2PService) AttachEndUserAuthorizedRep1ToProfile(data EndUserAssignmentData) (string, error) {
	params := &trusthub.CreateCustomerProfileEntityAssignmentParams{}
	params.SetObjectSid(data.EndUserSid)

	resp, err := s.client.TrusthubV1.CreateCustomerProfileEntityAssignment(data.CustomerProfileSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to attach EndUser authorized representative 1 to customer profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.6 Create An Address Resource and returns address sid
func (s *A2PService) CreateAddressResource(data AddressData) (string, error) {
	params := &api.CreateAddressParams{}
	// params.SetPathAccountSid(data.PathAccountSid)
	params.SetCustomerName(data.CustomerName)
	params.SetStreet(data.Street)
	params.SetCity(data.City)
	params.SetRegion(data.Region)
	params.SetPostalCode(data.PostalCode)
	params.SetIsoCountry(data.IsoCountry)
	params.SetFriendlyName(data.FriendlyName)
	params.SetStreetSecondary(data.StreetSecondary)
	params.SetAutoCorrectAddress(true)

	resp, err := s.client.Api.CreateAddress(params)
	if err != nil {
		return "", fmt.Errorf("failed to create Address: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.7 Create a supporting document resource and returns supporting_document_sid
func (s *A2PService) CreateSupportingDocumentResource(data SupportingDocumentData) (string, error) {
	params := &trusthub.CreateSupportingDocumentParams{}
	params.SetFriendlyName(data.FriendlyName)
	params.SetType("customer_profile_address")
	params.SetAttributes(map[string]interface{}{
		"address_sids": data.AddressSid,
	})

	resp, err := s.client.TrusthubV1.CreateSupportingDocument(params)
	if err != nil {
		return "", fmt.Errorf("failed to create Supporting Document: %w", err)
	}
	return *resp.Sid, nil
}

// Step 2.8 Attach the SupportingDocument resource to the Secondary Customer Profile
func (s *A2PService) AttachSupportingDocumentToProfile(secondaryProfileSID string, supportingDocumentSID *string) (string, error) {
	params := &trusthub.CreateCustomerProfileEntityAssignmentParams{}
	params.SetObjectSid(*supportingDocumentSID)

	resp, err := s.client.TrusthubV1.CreateCustomerProfileEntityAssignment(secondaryProfileSID, params)
	if err != nil {
		return "", fmt.Errorf("failed to attach Supporting Document to customer profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.9. Evaluate the Secondary Customer Profile
func (s *A2PService) EvaluateSecondaryCustomerProfile(secondaryProfileSID string) (string, error) {
	params := &trusthub.CreateCustomerProfileEvaluationParams{}
	params.SetPolicySid("RNdfbf3fae0e1107f8aded0e7cead80bf5")

	resp, err := s.client.TrusthubV1.CreateCustomerProfileEvaluation(secondaryProfileSID, params)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate Secondary Customer Profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 2.10. Submit the Secondary Customer Profile for review  - status must be set to pending-review
func (s *A2PService) SubmitSecondaryCustomerProfileForReview(secondaryProfileSID string) (string, error) {
	params := &trusthub.UpdateCustomerProfileParams{}
	params.SetStatus("pending-review")

	resp, err := s.client.TrusthubV1.UpdateCustomerProfile(secondaryProfileSID, params)
	if err != nil {
		return "", fmt.Errorf("failed to submit Secondary Customer Profile for review: %w", err)
	}

	return *resp.Sid, nil
}
