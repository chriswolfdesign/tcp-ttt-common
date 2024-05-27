package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tcp-ttt-common/model"
	"tcp-ttt-common/enums"
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

	Describe("get winner tests", func() {
		var board model.Board

		BeforeEach(func() {
			board = model.GenerateBoard()
		})

		When("player one has won by the first row", func() {
			BeforeEach(func() {
				board.MakeMove(0, 0, enums.PLAYER_ONE)
				board.MakeMove(0, 1, enums.PLAYER_ONE)
				board.MakeMove(0, 2, enums.PLAYER_ONE)
			})

			It("Returns player one as the winner", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two has won by the first column", func() {
			BeforeEach(func() {
				board.MakeMove(0, 0, enums.PLAYER_TWO)
				board.MakeMove(1, 0, enums.PLAYER_TWO)
				board.MakeMove(2, 0, enums.PLAYER_TWO)
			})

			It("returns player two as the winner", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(enums.PLAYER_TWO))
			})
		})

		When("player one has won by the top-left to bottom-right diagonal", func() {
			BeforeEach(func() {
				board.MakeMove(0, 0, enums.PLAYER_ONE)
				board.MakeMove(1, 1, enums.PLAYER_ONE)
				board.MakeMove(2, 2, enums.PLAYER_ONE)
			})

			It("returns player one as the winner", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two has won by the top-right to bottom-left diagonal", func() {
			BeforeEach(func() {
				board.MakeMove(2, 0, enums.PLAYER_TWO)
				board.MakeMove(1, 1, enums.PLAYER_TWO)
				board.MakeMove(0, 2, enums.PLAYER_TWO)
			})

			It("returns player two as the winner", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(enums.PLAYER_TWO))
			})
		})

		When("the game is not over", func() {
			It("returns that no player has won the game yet", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(strings.NOT_OVER))
			})
		})

		When("nobody has won but there is no valid move left", func() {
			BeforeEach(func() {
				board.MakeMove(0, 0, enums.PLAYER_ONE)
				board.MakeMove(0, 1, enums.PLAYER_TWO)
				board.MakeMove(0, 2, enums.PLAYER_ONE)
	
				board.MakeMove(1, 0, enums.PLAYER_TWO)
				board.MakeMove(1, 1, enums.PLAYER_ONE)
				board.MakeMove(1, 2, enums.PLAYER_TWO)
	
				board.MakeMove(2, 0, enums.PLAYER_TWO)
				board.MakeMove(2, 1, enums.PLAYER_ONE)
				board.MakeMove(2, 2, enums.PLAYER_TWO)
			})

			It("returns the game was a draw", func() {
				winner := board.GetWinner()

				Expect(winner).To(Equal(strings.DRAW))
			})
		})
	})
})
