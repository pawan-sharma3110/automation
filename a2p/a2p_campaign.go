// Step 5: Create an A2P Campaign
package a2p

import (
	"fmt"

	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

// Step 6.1: FetchA2PUseCases fetches the possible A2P campaign use cases for a given brand registration
// Note : Do not complete this section until the BrandRegistration's status is APPROVED.
// TODO: need to add logic at the time of creating subaccount to check the status of the BrandRegistration
func (s *A2PService) FetchA2PUseCases(messagingServiceSid, brandRegistrationSid string) ([]messaging.MessagingV1UsAppToPersonUsecase, error) {
	params := &messaging.FetchUsAppToPersonUsecaseParams{}
	params.SetBrandRegistrationSid(brandRegistrationSid)

	resp, err := s.client.MessagingV1.FetchUsAppToPersonUsecase(messagingServiceSid, params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch A2P use cases: %w", err)
	}

	useCases := []messaging.MessagingV1UsAppToPersonUsecase{}
	for _, useCase := range *resp.UsAppToPersonUsecases {
		parsedUseCase, ok := useCase.(messaging.MessagingV1UsAppToPersonUsecase)
		if !ok {
			return nil, fmt.Errorf("failed to parse use case")
		}
		useCases = append(useCases, parsedUseCase)
	}

	return useCases, nil
}

// Step 6.2: Create the A2P Campaign
// Note : Do not complete this section until the BrandRegistration's status is APPROVED.
// TODO: need to add logic at the time of creating subaccount to check the status of the BrandRegistration
func (s *A2PService) CreateA2PCampaign(messagingServiceSid string, data CampaignData) (string, error) {
	params := &messaging.CreateUsAppToPersonParams{}
	// params.SetDescription(data.Description)
	// params.SetUsAppToPersonUsecase(data.Usecase)
	// params.SetHasEmbeddedLinks(data.HasEmbeddedLinks)
	// params.SetHasEmbeddedPhone(data.HasEmbeddedPhone)
	// params.SetMessageSamples(data.MessageSamples)
	// params.SetMessageFlow(data.MessageFlow)
	// params.SetBrandRegistrationSid(data.BrandRegistrationSid)

	params.SetUsAppToPersonUsecase("DELIVERY_NOTIFICATION")
	params.SetDescription("Delivery Notification Campaign")
	params.SetHasEmbeddedLinks(true)
	params.SetHasEmbeddedPhone(true)
	params.SetMessageSamples([]string{"Your appointment diagnosis for [disease] at [hospital_name] has been booked at [timestamp]. Please reply with 'YES' to confirm. If you need to reschedule, please reply with 'NO'. If you need any further assistance please call us at [phone_number] between [time]am to [time]pm from Monday to Friday. Thank you."})
	params.SetMessageFlow("Your appointment diagnosis for [disease] at [hospital_name] has been booked at [timestamp]. Please reply with 'YES' to confirm. If you need to reschedule, please reply with 'NO'. If you need any further assistance please call us at [phone_number] between [time]am to [time]pm from Monday to Friday. Thank you.")
	params.SetBrandRegistrationSid(data.BrandRegistrationSid)
	params.SetHelpKeywords([]string{"appointment", "diagnosis", "hospital", "booked", "timestamp", "confirm"})
	params.SetHelpMessage("Your appointment diagnosis for [disease] at [hospital_name] has been booked at [timestamp]. Please reply with 'YES' to confirm. If you need to reschedule, please reply with 'NO'. If you need any further assistance please call us at [phone_number] between [time]am to [time]pm from Monday to Friday. Thank you.")
	params.SetOptInKeywords([]string{"YES", "NO"})
	params.SetOptInMessage("Your appointment diagnosis for [disease] at [hospital_name] has been booked at [timestamp]. Please reply with 'YES' to confirm. If you need to reschedule, please reply with 'NO'. If you need any further assistance please call us at [phone_number] between [time]am to [time]pm from Monday to Friday. Thank you.")
	params.SetOptOutKeywords([]string{"STOP", "CANCEL"})
	params.SetOptOutMessage("Your appointment diagnosis for [disease] at [hospital_name] has been booked at [timestamp]. Please reply with 'YES' to confirm. If you need to reschedule, please reply with 'NO'. If you need any further assistance please call us at [phone_number] between [time]am to [time]pm from Monday to Friday. Thank you.")

	resp, err := s.client.MessagingV1.CreateUsAppToPerson(messagingServiceSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to create A2P Campaign: %w", err)
	}

	return *resp.Sid, nil
}

// Step 10.2: View and Manage A2P Campaigns
func (s *A2PService) ListA2PCampaigns(messagingServiceSid string) ([]CampaignData, error) {
	params := &messaging.ListUsAppToPersonParams{}
	campaigns, err := s.client.MessagingV1.ListUsAppToPerson(messagingServiceSid, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list A2P campaigns: %w", err)
	}

	var campaignList []CampaignData
	for _, campaign := range campaigns {
		campaignList = append(campaignList, CampaignData{
			SID:                  *campaign.Sid,
			Description:          *campaign.Description,
			Usecase:              *campaign.UsAppToPersonUsecase,
			CampaignStatus:       *campaign.CampaignStatus,
			BrandRegistrationSid: *campaign.BrandRegistrationSid,
		})
	}

	return campaignList, nil
}
