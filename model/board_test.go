package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tcp-ttt-common/model"
	"tcp-ttt-common/model/enums"
	"tcp-ttt-common/strings"
)

var _ = Describe("Board", func() {
	Describe("generator test", func() {
		When("Board is generated", func() {
			It("sets all squares to 0", func() {
				board := model.GenerateBoard()

				for i := range 3 {
					for j := range 3 {
						Expect(board.Board[i][j]).To(Equal(enums.EMPTY))
					}
				}
			})
		})
	})

	Describe("Make move tests", func() {
		var board model.Board

		BeforeEach(func() {
			board = model.GenerateBoard()
		})

		When("invalid string is provided", func() {
			It("throws appropriate error", func() {
				err := board.MakeMove(0, 0, "Y")

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.INCORRECT_PLAYER_VALUE))
			})
		})

		When("player one successfully moves", func() {
			It("sets the square correctly without throwing an error", func() {
				err := board.MakeMove(0, 0, enums.PLAYER_ONE)

				Expect(err).ToNot(HaveOccurred())
				Expect(board.Board[0][0]).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two successfully moves", func() {
			It("sets the square correctly without throwing an error", func() {
				err := board.MakeMove(2, 2, enums.PLAYER_TWO)

				Expect(err).ToNot(HaveOccurred())
				Expect(board.Board[2][2]).To(Equal(enums.PLAYER_TWO))
			})
		})

		When("player makes a move that is out of bounds", func() {
			It("throws an error", func() {
				err := board.MakeMove(3, 3, enums.PLAYER_ONE)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.MOVE_OUT_OF_BOUNDS))
			})
		})

		When("player makes a move on a nonempty square", func() {
			It("throws an error", func() {
				err := board.MakeMove(1, 1, enums.PLAYER_ONE)

				Expect(err).ToNot(HaveOccurred())
				Expect(board.Board[1][1]).To(Equal(enums.PLAYER_ONE))

				err = board.MakeMove(1, 1, enums.PLAYER_TWO)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.NOT_EMPTY_SQUARE))
				Expect(board.Board[1][1]).To(Equal(enums.PLAYER_ONE))
			})
		})
	})
})
