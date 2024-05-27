package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tcp-ttt-common/model"
)

var _ = Describe("Board", func() {
	Describe("generator test", func() {
		When("Board is generated", func() {
			It("sets all squares to 0", func() {
				board := model.GenerateBoard()

				for i := range 3 {
					for j := range 3 {
						Expect(board.Board[i][j]).To(Equal(0))
					}
				}
			})
		})
	})
})
