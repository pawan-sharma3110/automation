// // Step 7: Submit the A2P Campaign

package a2p

import (
	"fmt"
)

// Step 8.1: UpdateA2PCampaign updates the details of an A2P campaign
// func (s *A2PService) UpdateA2PCampaign(messagingServiceSid string, campaignSid string, params *messaging.upda) (string, error) {
// 	resp, err := s.client.UpdateUsAppToPerson(messagingServiceSid, campaignSid, params)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to update A2P Campaign: %w", err)
// 	}

// 	return *resp.CampaignStatus, nil
// }

// Step 8.2: CheckA2PCampaignStatus checks the status of an A2P campaign(Optional)

func (s *A2PService) CheckA2PCampaignStatus(messagingServiceSid string, campaignSid string) (string, error) {
	resp, err := s.client.MessagingV1.FetchUsAppToPerson(messagingServiceSid, campaignSid)
	if err != nil {
		return "", fmt.Errorf("failed to check A2P Campaign status: %w", err)
	}

	return *resp.CampaignStatus, nil
}
