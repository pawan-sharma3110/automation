package a2p

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

type PhoneNumberData struct {
	PhoneNumberSid string
}

type ListAvailablePhoneNumberLocalParams struct {
	AreaCode     int
	Contains     string
	InRegion     string
	SmsEnabled   bool
	VoiceEnabled bool
}

// GetAvailablePhoneNumbers retrieves a list of available local phone numbers
// for a specified area code and prints their friendly names.
func (s *A2PService) GetAvailablePhoneNumbers(param ListAvailablePhoneNumberLocalParams) ([]string, error) {
	client := twilio.NewRestClient()

	// Set parameters for the API request
	params := &api.ListAvailablePhoneNumberLocalParams{
		AreaCode:     &param.AreaCode,
		Contains:     &param.Contains,
		SmsEnabled:   &param.SmsEnabled,
		VoiceEnabled: &param.VoiceEnabled,
		InRegion:     &param.InRegion,
	}

	// Make the API request to fetch available phone numbers
	resp, err := client.Api.ListAvailablePhoneNumberLocal("US", params)
	if err != nil {
		return nil, fmt.Errorf("error fetching phone numbers: %v", err)
	}

	// Collect the friendly names of the available phone numbers
	var phoneNumbers []string
	for _, record := range resp {
		if record.FriendlyName != nil {
			phoneNumbers = append(phoneNumbers, *record.FriendlyName)
		} else {
			phoneNumbers = append(phoneNumbers, "No friendly name available")
		}
	}
	return phoneNumbers, nil
}

// AddPhoneNumberToMessagingService associates a phone number with the messaging service.
func (s *A2PService) AddPhoneNumberToMessagingService(serviceSid string, params *messaging.CreatePhoneNumberParams) (*messaging.MessagingV1PhoneNumber, error) {
	// Add the phone number to the Messaging Service
	resp, err := s.client.MessagingV1.CreatePhoneNumber(serviceSid, params)
	if err != nil {
		return nil, fmt.Errorf("error adding phone number to messaging service: %v", err)
	}
	if resp.Sid != nil {
		fmt.Println("Phone number added successfully with SID:", *resp.Sid)
	} else {
		fmt.Println("Failed to retrieve SID from response.")
	}
	return resp, nil
}

func (s *A2PService) GetPhoneNumberSID(phoneNumber string) (string, error) {

	params := &api.ListIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := s.client.Api.ListIncomingPhoneNumber(params)
	if err != nil {
		return "", err
	}

	if len(resp) == 0 {
		return "", fmt.Errorf("no phone number found for %s", phoneNumber)
	}

	return *resp[0].Sid, nil
}
