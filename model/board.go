package model

type Board struct {
	Board [][]int
}

func GenerateBoard() Board {
	tmp := make([][]int, 3)
	for i := range 3 {
		tmp[i] = make([]int, 3)
	}

	return Board{
		Board: tmp,
	}
}