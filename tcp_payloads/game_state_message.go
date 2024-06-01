package tcp_payloads

import "github.com/chriswolfdesign/tcp-ttt-common/model"

type GameStateMessage struct {
	Game model.Game
	PayloadType string
}
