package model

import (
	"fmt"
	"tcp-ttt-common/enums"
	"tcp-ttt-common/strings"
)

type Game struct {
	Board Board
	CurrentPlayer string
	PlayerOneName string
	PlayerTwoName string
	Winner string
}

func GenerateGame() *Game {
	return &Game{
		Board: GenerateBoard(),
		CurrentPlayer: enums.PLAYER_ONE,
		Winner: strings.NOT_OVER,
	}
}

func (g *Game) SetPlayerOne(name string) error {
	if g.PlayerOneName != "" {
		return fmt.Errorf(strings.PLAYER_ONE_ALREADY_SET)
	}

	g.PlayerOneName = name

	return nil
}

func (g *Game) SetPlayerTwo(name string) error {
	if g.PlayerTwoName != "" {
		return fmt.Errorf(strings.PLAYER_TWO_ALREADY_SET)
	}

	g.PlayerTwoName = name

	return nil
}

func (g *Game) MakeMove(row, col int) error {
	err := g.Board.MakeMove(row, col, g.CurrentPlayer)

	if err != nil {
		return err
	}

	if (g.CurrentPlayer == enums.PLAYER_ONE) {
		g.CurrentPlayer = enums.PLAYER_TWO
	} else {
		g.CurrentPlayer = enums.PLAYER_ONE
	}

	g.Winner = g.Board.GetWinner()

	return nil
}