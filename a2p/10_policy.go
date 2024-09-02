// Step 11: Handling PolicySID
package a2p

import (
	openapi "github.com/twilio/twilio-go/rest/trusthub/v1"
)

type PolicyInfo struct {
	Sid          *string      `json:"sid,omitempty"`
	FriendlyName *string      `json:"friendly_name,omitempty"`
	Requirements *interface{} `json:"requirements,omitempty"`
	Url          *string      `json:"url,omitempty"`
}

// Step 10.1: Fetch Available Policies
// PolicySId already set to hard coded value, so this function is not needed. (optional)
// Do not include this function in the final code.
func (s *A2PService) ListPolicies(sid string, pageSize, limit *int) (*openapi.TrusthubV1Policies, error) {

	policies, err := s.client.TrusthubV1.FetchPolicies(sid)
	if err != nil {
		return nil, err
	}

	return policies, nil
}
