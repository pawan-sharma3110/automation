// Step 2: Create and Manage a TrustProduct
package a2p

import (
	"fmt"

	trusthub "github.com/twilio/twilio-go/rest/trusthub/v1"
)

// Step 3.1: Create a TrustProduct Resource
func (s *A2PService) CreateTrustProduct(data TrustProductData) (string, error) {
	params := &trusthub.CreateTrustProductParams{}
	params.SetFriendlyName(data.FriendlyName)
	params.SetPolicySid(data.PolicySid)
	params.SetEmail(data.Email)
	if data.StatusCallback != "" {
		params.SetStatusCallback(data.StatusCallback)
	}

	resp, err := s.client.TrusthubV1.CreateTrustProduct(params)
	if err != nil {
		return "", fmt.Errorf("failed to create TrustProduct: %w", err)
	}

	return *resp.Sid, nil
}

// Step 3.2: Create an EndUser Resource of Type us_a2p_messaging_profile_information
// Note :  a2p_messaging_profile_sid will be the id returned from this function
func (s *A2PService) CreateEndUserMessagingProfile(data EndUserMessagingProfileData) (string, error) {
	params := &trusthub.CreateEndUserParams{}
	params.SetAttributes(map[string]interface{}{
		"company_type":   data.CompanyType,
		"stock_exchange": data.StockExchange,
		"stock_ticker":   data.StockTicker,
	})
	params.SetFriendlyName(fmt.Sprintf("%s Messaging Profile EndUser", data.CompanyType))
	params.SetType("us_a2p_messaging_profile_information")

	resp, err := s.client.TrusthubV1.CreateEndUser(params)
	if err != nil {
		return "", fmt.Errorf("failed to create EndUser messaging profile: %w", err)
	}

	return *resp.Sid, nil
}

// Step 3.3: Attach the EndUser to the TrustProduct
func (s *A2PService) AttachEndUserToTrustProduct(trustProductSid, endUserSid string) (string, error) {
	params := &trusthub.CreateTrustProductEntityAssignmentParams{}
	params.SetObjectSid(endUserSid)

	resp, err := s.client.TrusthubV1.CreateTrustProductEntityAssignment(trustProductSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to attach EndUser to TrustProduct: %w", err)
	}

	return *resp.Sid, nil
}

// Step 3.4: Attach the Secondary Customer Profile to the TrustProduct
func (s *A2PService) AttachSecondaryCustomerProfileToTrustProduct(trustProductSid, customerProfileSid string) (string, error) {
	params := &trusthub.CreateTrustProductEntityAssignmentParams{}
	params.SetObjectSid(customerProfileSid)

	resp, err := s.client.TrusthubV1.CreateTrustProductEntityAssignment(trustProductSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to attach Customer Profile to TrustProduct: %w", err)
	}

	return *resp.Sid, nil
}

// Step 3.5: Evaluate the TrustProduct
func (s *A2PService) EvaluateTrustProduct(trustProductSid, policySid string) (string, error) {
	params := &trusthub.CreateTrustProductEvaluationParams{}
	params.SetPolicySid(policySid)

	resp, err := s.client.TrusthubV1.CreateTrustProductEvaluation(trustProductSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate TrustProduct: %w", err)
	}

	return *resp.Sid, nil
}

// Step 3.6: Submit the TrustProduct for Review  - status must be set to pending-review
func (s *A2PService) SubmitTrustProductForReview(trustProductSid string) (string, error) {
	params := &trusthub.UpdateTrustProductParams{}
	params.SetStatus("pending-review")

	resp, err := s.client.TrusthubV1.UpdateTrustProduct(trustProductSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to submit TrustProduct for review: %w", err)
	}

	return *resp.Sid, nil
}
