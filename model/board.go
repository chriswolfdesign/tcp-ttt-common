package model

import "tcp-ttt-common/model/enums"

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