package tcp_payloads

import "github.com/chriswolfdesign/tcp-ttt-common/strings"

type PlayerOnboardingRequest struct {
	Name string
	PayloadType string
}

func GeneratePlayerOnboardingRequest(name string) PlayerOnboardingRequest {
	return PlayerOnboardingRequest{
		Name: name,
		PayloadType: strings.TYPE_ONBOARDING_REQUEST,
	}
}