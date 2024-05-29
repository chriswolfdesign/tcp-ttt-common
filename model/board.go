package model

import (
	"fmt"
	"github.com/chriswolfdesign/tcp-ttt-common/enums"
	"github.com/chriswolfdesign/tcp-ttt-common/strings"
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

func (b Board) PrintBoard() {
	fmt.Println("   0   1   2")
	for i := range 3 {
		fmt.Printf("%d ", i)
		for j := range 3 {
			fmt.Printf("[%s] ", b.Board[i][j])
		}
		fmt.Println()
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

func (b Board) GetWinner() string {
	if b.hasPlayerWon(enums.PLAYER_ONE) {
		return enums.PLAYER_ONE
	}

	if b.hasPlayerWon(enums.PLAYER_TWO) {
		return enums.PLAYER_TWO
	}

	if b.emptySquares() == 0 {
		return strings.DRAW
	}

	return strings.NOT_OVER
}

func (b Board) hasPlayerWon(player string) bool {
	return b.hasPlayerWonByRow(player) || b.hasPlayerWonByCol(player) || b.hasPlayerWonByDiagonals(player)
}

func (b Board) hasPlayerWonByRow(player string) bool {
	for i := range 3 {
		if b.Board[i][0] == player && b.Board[i][1] == player && b.Board[i][2] == player {
			return true
		}
	}

	return false
}

func (b Board) hasPlayerWonByCol(player string) bool {
	for i := range 3 {
		if b.Board[0][i] == player && b.Board[1][i] == player && b.Board[2][i] == player {
			return true
		}
	}

	return false
}

func (b Board) hasPlayerWonByDiagonals(player string) bool {
	if b.Board[0][0] == player && b.Board[1][1] == player && b.Board[2][2] == player {
		return true
	}

	if b.Board[2][0] == player && b.Board[1][1] == player && b.Board[0][2] == player {
		return true
	}

	return false
}

func (b Board) emptySquares() int {
	res := 0

	for i := range 3 {
		for j := range 3 {
			if b.Board[i][j] == enums.EMPTY {
				res += 1
			}
		}
	}

	return res
}

