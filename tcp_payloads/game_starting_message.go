package tcp_payloads

import "github.com/chriswolfdesign/tcp-ttt-common/model"

type GameStartingMessage struct {
	Message     string
	PayloadType string
	Game        model.Game
}
