// Step 9: Automating Customer Onboarding
package a2p

import (
	"fmt"

	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SubAccountData struct {
	FriendlyName string
}

// Step 1.0: Automate Subaccount Creation
func (s *A2PService) CreateSubAccount(data SubAccountData) (string, string, error) {
	params := &api.CreateAccountParams{}
	params.SetFriendlyName(data.FriendlyName)

	resp, err := s.client.Api.CreateAccount(params)
	if err != nil {
		return "", "", fmt.Errorf("failed to create subaccount: %w", err)
	}

	return *resp.Sid, *resp.AuthToken, nil
}
