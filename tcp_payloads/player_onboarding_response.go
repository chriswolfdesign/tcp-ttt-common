package tcp_payloads

import "github.com/chriswolfdesign/tcp-ttt-common/strings"

type PlayerOnboardingResponse struct {
	Status      string
	Player      string
	PayloadType string
}

func GeneratePlayerOnboardingResponse(status, player string) PlayerOnboardingResponse {
	return PlayerOnboardingResponse{
		Status:      status,
		Player:      player,
		PayloadType: strings.TYPE_ONBOARDING_RESPONSE,
	}
}

