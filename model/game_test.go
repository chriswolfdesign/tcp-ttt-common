package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tcp-ttt-common/enums"
	"tcp-ttt-common/model"
	"tcp-ttt-common/strings"
)

var _ = Describe("Game", func() {
	Describe("Generator tests", func() {
		var game model.Game

		BeforeEach(func() {
			game = *model.GenerateGame()
		})

		When("The game is generated", func() {
			It("initializes everything correctly", func() {
				for i := range 3 {
					for j := range 3 {
						Expect(game.Board.Board[i][j]).To(Equal(enums.EMPTY))
					}
				}

				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_ONE))
				Expect(game.Winner).To(Equal(strings.NOT_OVER))
				Expect(game.PlayerOneName).To(Equal(""))
				Expect(game.PlayerTwoName).To(Equal(""))
			})
		})
	})

	Describe("set player one tests", func() {
		var game model.Game

		BeforeEach(func() {
			game = *model.GenerateGame()
		})

		When("player one has not yet been set", func() {
			It("sets player one's name correctly", func() {
				err := game.SetPlayerOne("Chris")

				Expect(err).ToNot(HaveOccurred())
				Expect(game.PlayerOneName).To(Equal("Chris"))
			})
		})

		When("player one has already been set", func() {
			BeforeEach(func() {
				err := game.SetPlayerOne("Chris")
				Expect(err).ToNot(HaveOccurred())
				Expect(game.PlayerOneName).To(Equal("Chris"))
			})

			It("throws an error and does not update the name", func() {
				err := game.SetPlayerOne("Wolf")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.PLAYER_ONE_ALREADY_SET))
				Expect(game.PlayerOneName).To(Equal("Chris"))
			})
		})
	})

	Describe("set player two tests", func() {
		var game model.Game

		BeforeEach(func() {
			game = *model.GenerateGame()
		})

		When("palyer two has not yet been set", func() {
			It("sets player two's name correctly", func() {
				err := game.SetPlayerTwo("Chris")
				Expect(err).ToNot(HaveOccurred())
				Expect(game.PlayerTwoName).To(Equal("Chris"))
			})
		})

		When("player two has already been set", func() {
			BeforeEach(func() {
				err := game.SetPlayerTwo("Chris")
				Expect(err).ToNot(HaveOccurred())
				Expect(game.PlayerTwoName).To(Equal("Chris"))
			})

			It("throws and error and does not update the name", func() {
				err := game.SetPlayerTwo("Wolf")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.PLAYER_TWO_ALREADY_SET))
				Expect(game.PlayerTwoName).To(Equal("Chris"))
			})
		})
	})

	Describe("make move tests", func() {
		var game model.Game

		BeforeEach(func() {
			game = *model.GenerateGame()
		})

		When("Player one makes a legal move", func() {
			It("updates the board and switches the current player", func() {
				err := game.MakeMove(1, 1)

				Expect(err).ToNot(HaveOccurred())
				Expect(game.Board.Board[1][1]).To(Equal(enums.PLAYER_ONE))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_TWO))
			})
		})

		When("player one makes an illegal move", func() {
			It("does not update the board and does switch the current player", func() {
				err := game.MakeMove(-1, 0)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.MOVE_OUT_OF_BOUNDS))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two makes a legal move", func() {
			BeforeEach(func() {
				err := game.MakeMove(0, 0)

				Expect(err).ToNot(HaveOccurred())
				Expect(game.Board.Board[0][0]).To(Equal(enums.PLAYER_ONE))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_TWO))
			})

			It("updates the board and switches the current player", func() {
				err := game.MakeMove(1, 1)

				Expect(err).ToNot(HaveOccurred())
				Expect(game.Board.Board[0][0]).To(Equal(enums.PLAYER_ONE))
				Expect(game.Board.Board[1][1]).To(Equal(enums.PLAYER_TWO))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two makes an illegal move", func() {
			BeforeEach(func() {
				err := game.MakeMove(0, 0)

				Expect(err).ToNot(HaveOccurred())
				Expect(game.Board.Board[0][0]).To(Equal(enums.PLAYER_ONE))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_TWO))
			})

			It("does not update the board and does not switch the current player", func() {
				err := game.MakeMove(0, 0)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(strings.NOT_EMPTY_SQUARE))
				Expect(game.CurrentPlayer).To(Equal(enums.PLAYER_TWO))
			})
		})

		When("player one wins the game", func() {
			BeforeEach(func() {
				game.MakeMove(0, 0)
				game.MakeMove(1, 0)
				game.MakeMove(0, 1)
				game.MakeMove(2, 0)
				game.MakeMove(0, 2)
			})

			It("updates the game state accordingly", func() {
				Expect(game.Winner).To(Equal(enums.PLAYER_ONE))
			})
		})

		When("player two winds the game", func() {
			BeforeEach(func() {
				game.MakeMove(0, 0)
				game.MakeMove(1, 0)
				game.MakeMove(0, 1)
				game.MakeMove(1, 1)
				game.MakeMove(2, 0)
				game.MakeMove(1, 2)
			})

			It("updates the game state accordingly", func() {
				Expect(game.Winner).To(Equal(enums.PLAYER_TWO))
			})
		})
		
		When("the game ends in a draw", func() {
			BeforeEach(func() {
				game.MakeMove(0, 0)
				game.MakeMove(0, 1)
				game.MakeMove(0, 2)
				game.MakeMove(1, 0)
				game.MakeMove(2, 0)
				game.MakeMove(1, 1)
				game.MakeMove(2, 1)
				game.MakeMove(2, 2)
				game.MakeMove(1, 2)
			})

			It("updates the game state accordingly", func() {
				Expect(game.Winner).To(Equal(strings.DRAW))
			})
		})
	})
})
