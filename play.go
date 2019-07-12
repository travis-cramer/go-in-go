package main

import (
	"fmt"
)

func (game *Game) play() {
	game.printGame()

	game.gameOn = true
	for game.gameOn == true {
		piecePlaced := false
		playerPassed := false
		for piecePlaced == false {
			row := 0
			for row == 0 {
				row = promptForInt("Enter row: ")
			}
			col := 0
			for col == 0 {
				col = promptForInt("Enter column: ")
			}
			playerPassed = game.checkForPass(row, col, playerPassed)
			if (playerPassed == false) {
				piecePlaced = game.placePiece(row, col, game.currentPlayer)
			}
			if game.gameOn == false {
				break
			}
		}
		if game.gameOn == false {
			break
		}
		game.changeTurn()
		game.printGame()
	}
	game.gameOver(game.board)
}

func (game *Game) placePiece(col int, row int, player Player) bool {
	if !game.indexIsOnBoard(row-1, col-1) {
		fmt.Printf("Off the game board.\nTry again.\n")
		return false
	}

	game.board[row-1][col-1] = player.toInt()

	piecesCaptured := game.checkForAndRemoveCapturedPieces()
	if !piecesCaptured {
		selfCapture := game.checkForAndRemoveSelfCapturedPieces()
		if selfCapture {
			fmt.Printf("Invalid move.\nTry again.\n")
			return false
		}
	}
	return true
}

func (game *Game) changeTurn() {
	if game.currentPlayer == LIGHT {
		game.currentPlayer = DARK
	} else if game.currentPlayer == DARK {
		game.currentPlayer = LIGHT
	}
}

func (game *Game) indexIsOnBoard(i int, j int) bool {
	if i >= game.boardSize || i < 0 {
		return false
	}
	if j >= game.boardSize || j < 0 {
		return false
	}
	return true
}

func (game *Game) checkForPass(row int, col int, passedOnce bool) bool {
	if (passedOnce) {
		if (row == -1 && col == -1) {
			fmt.Printf("%s player has passed. The game is now over, and the tallying process will begin. \n", game.currentPlayer.toString())
			game.gameOn = false
			return true
		}
	} else {
		if (row == -1 && col == -1) {
			fmt.Printf("%s player has passed, %s player's turn. \n", game.currentPlayer.toString(), game.currentPlayer.opposingPlayer().toString())
			return true
		}
	}	

	return false
}