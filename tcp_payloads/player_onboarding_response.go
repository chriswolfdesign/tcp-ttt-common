package tcp_payloads

import "github.com/chriswolfdesign/tcp-ttt-common/strings"

type PlayerOnboardingResponse struct {
	Status string
	PayloadType string
}

func GeneratePlayerOnboardingResponse(status string) PlayerOnboardingResponse {
	return PlayerOnboardingResponse{
		Status: status,
		PayloadType: strings.TYPE_ONBOARDING_RESPONSE,
	}
}