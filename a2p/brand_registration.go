// Step 3: Create a BrandRegistration
// We'll start by creating the BrandRegistration resource. This will represent your customer's brand and submit the Brand-related information for vetting.
package a2p

import (
	"fmt"

	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

// Step 4.1: Create a BrandRegistration
/*
Note :
The customer_profile_bundle_sid is the SID of your customer's Secondary Customer Profile.
The a2p_profile_bundle_sid is the SID of the TrustProduct created SID.
Skip_automatic_sec_vet is an optional Boolean.
Sometimes, Brand vetting by TCR can take several days.
If the BrandRegistration resources's status is IN_REVIEW for more than two days then please contact to the Twilio Support.
*/
func (s *A2PService) CreateBrandRegistration(data BrandRegistrationData) (string, string, error) {
	params := &messaging.CreateBrandRegistrationsParams{}
	params.SetCustomerProfileBundleSid(data.CustomerProfileBundleSid)
	params.SetA2PProfileBundleSid(data.A2PProfileBundleSid)

	resp, err := s.client.MessagingV1.CreateBrandRegistrations(params)
	if err != nil {
		return "", "", fmt.Errorf("failed to create BrandRegistration: %w", err)
	}
	return *resp.Sid, *resp.Status, nil
}

// Step 4.2: Skipping Secondary Vetting
// (Optional) Do Not Use This Function If You Are Not Sure
func (s *A2PService) CreateBrandRegistrationWithSkipVetting(data BrandRegistrationData, skipVetting bool) (string, error) {
	params := &messaging.CreateBrandRegistrationsParams{}
	params.SetCustomerProfileBundleSid(data.CustomerProfileBundleSid)
	params.SetA2PProfileBundleSid(data.A2PProfileBundleSid)
	params.SetSkipAutomaticSecVet(skipVetting)

	resp, err := s.client.MessagingV1.CreateBrandRegistrations(params)
	if err != nil {
		return "", fmt.Errorf("failed to create BrandRegistration with skip vetting: %w", err)
	}

	return *resp.Sid, nil
}

func (s *A2PService) FetchBrandRegistration(sid string) (string, error) {
	resp, err := s.client.MessagingV1.FetchBrandRegistrations(sid)
	if err != nil {
		return "", fmt.Errorf("failed to fetch BrandRegistration: %w", err)
	}
	return *resp.Status, nil
}

func (s *A2PService) ListBrandRegistrations() ([]messaging.MessagingV1BrandRegistrations, error) {
	resp, err := s.client.MessagingV1.ListBrandRegistrations(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list BrandRegistrations: %w", err)
	}
	return resp, nil
}
