// Step 4: Create a MessagingService
package a2p

import (
	"fmt"

	_ "github.com/twilio/twilio-go"
	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

// Step 5.1: Create a MessagingService Resource - This will return MessageServiceSID
func (s *A2PService) CreateMessagingService(data MessagingServiceData) (string, error) {
	params := &messaging.CreateServiceParams{}
	params.SetFriendlyName(data.FriendlyName)
	params.SetInboundRequestUrl(data.InboundRequestUrl)
	params.SetFallbackUrl(data.FallbackUrl)

	resp, err := s.client.MessagingV1.CreateService(params)
	if err != nil {
		return "", fmt.Errorf("failed to create MessagingService: %w", err)
	}
	return *resp.Sid, nil
}

// Step 5.2: Additional Configuration (Optional)
func (s *A2PService) CreateMessagingServiceWithConfig(data MessagingServiceAdditional) (string, error) {
	params := &messaging.CreateServiceParams{}
	params.SetFriendlyName(data.FriendlyName)
	params.SetInboundRequestUrl(data.InboundRequestUrl)
	params.SetFallbackUrl(data.FallbackUrl)
	if data.StatusCallback != "" {
		params.SetStatusCallback(data.StatusCallback)
	}
	params.SetStickySender(data.StickySender)
	params.SetSmartEncoding(data.SmartEncoding)
	params.SetMmsConverter(data.MmsConverter)
	params.SetFallbackToLongCode(data.FallbackToLongCode)
	params.SetScanMessageContent(data.ScanMessageContent)
	params.SetAreaCodeGeomatch(data.AreaCodeGeomatch)
	params.SetValidityPeriod(data.ValidityPeriod)
	params.SetSynchronousValidation(data.SynchronousValidation)
	params.SetUsecase(data.Usecase)

	resp, err := s.client.MessagingV1.CreateService(params)
	if err != nil {
		return "", fmt.Errorf("failed to create MessagingService with config: %w", err)
	}

	return *resp.Sid, nil
}
