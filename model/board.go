package model

import (
	"fmt"
	"tcp-ttt-common/model/enums"
	"tcp-ttt-common/strings"
)

type Board struct {
	Board [][]string
}

func GenerateBoard() Board {
	tmp := make([][]string, 3)
	for i := range 3 {
		tmp[i] = make([]string, 3)
		for j := range 3 {
			tmp[i][j] = enums.EMPTY
		}
	}

	return Board{
		Board: tmp,
	}
}

func (b *Board) MakeMove(row int, col int, player string) error {
	if player != enums.PLAYER_ONE && player != enums.PLAYER_TWO {
		return fmt.Errorf(strings.INCORRECT_PLAYER_VALUE)
	}

	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return fmt.Errorf(strings.MOVE_OUT_OF_BOUNDS)
	}

	if b.Board[row][col] != enums.EMPTY {
		return fmt.Errorf(strings.NOT_EMPTY_SQUARE)
	}

	b.Board[row][col] = player

	return nil
}