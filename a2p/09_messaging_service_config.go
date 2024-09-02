// Step 8: Post-Approval Configurations and Monitoring

package a2p

import (
	"fmt"

	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

// Step 9.1: Finalize Messaging Service Configuration
func (s *A2PService) FinalizeMessagingServiceConfig(data FinalizeMessagingServiceConfigData) (string, error) {
	params := &messaging.UpdateServiceParams{}
	params.SetStatusCallback(data.StatusCallback)
	params.SetStickySender(data.StickySender)
	params.SetSmartEncoding(data.SmartEncoding)
	params.SetMmsConverter(data.MmsConverter)
	params.SetFallbackToLongCode(data.FallbackToLongCode)
	params.SetScanMessageContent(data.ScanMessageContent)
	params.SetAreaCodeGeomatch(data.AreaCodeGeomatch)
	params.SetValidityPeriod(data.ValidityPeriod)
	params.SetSynchronousValidation(data.SynchronousValidation)

	resp, err := s.client.MessagingV1.UpdateService(data.MessagingServiceSid, params)
	if err != nil {
		return "", fmt.Errorf("failed to finalize Messaging Service config: %w", err)
	}

	return *resp.Sid, nil
}
