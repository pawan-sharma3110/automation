// Step 10: Provide Customer Management and Reporting Tools

package a2p

import (
	"fmt"

	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SubAccountInfo struct {
	Sid          string
	FriendlyName string
	Status       string
}

// Step 10.1.1: View Subaccounts
func (s *A2PService) ListSubAccounts() ([]SubAccountInfo, error) {
	params := &api.ListAccountParams{}
	accounts, err := s.client.Api.ListAccount(params)
	if err != nil {
		return nil, fmt.Errorf("failed to list subaccounts: %w", err)
	}

	var subAccounts []SubAccountInfo
	for _, account := range accounts {
		subAccounts = append(subAccounts, SubAccountInfo{
			Sid:          *account.Sid,
			FriendlyName: *account.FriendlyName,
			Status:       *account.Status,
		})
	}

	return subAccounts, nil
}

// Step 10.1.2: Manage Subaccounts
func (s *A2PService) UpdateSubAccountStatus(subAccountSid, status string) (string, error) {
	params := &api.UpdateAccountParams{}
	params.SetStatus(status)

	resp, err := s.client.Api.UpdateAccount(subAccountSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to update subaccount status: %w", err)
	}

	return *resp.Sid, nil
}
