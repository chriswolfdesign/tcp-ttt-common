package tcp_payloads

type MakeMoveMessage struct {
	Player string
	Row int
	Col int
	PayloadType string
}