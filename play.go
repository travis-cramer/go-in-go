package main

import (
	"fmt"
)

func (game *Game) play() {
	game.printGame()

	game.gameOn = true
	for game.gameOn == true {
		piecePlaced := false
		for piecePlaced == false {
			row := 0
			for row == 0 {
				row = promptForInt("Enter row: ")
			}
			col := 0
			for col == 0 {
				col = promptForInt("Enter column: ")
			}
			piecePlaced = game.placePiece(row, col, game.currentPlayer)
		}
		game.changeTurn()
		game.printGame()
	}
}

func (game *Game) placePiece(row int, col int, player Player) bool {
	if !game.indexIsOnBoard(row-1, col-1) {
		fmt.Printf("Off the game board.\nTry again.\n")
		return false
	}

	game.board[row-1][col-1] = uint8(player)

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

func (game *Game) checkForAndRemoveCapturedPieces() bool {
	opposingPlayer := game.currentPlayer.opposingPlayer()
	return game.checkForAndRemoveCapturedPiecesOfPlayer(opposingPlayer)
}

func (game *Game) checkForAndRemoveSelfCapturedPieces() bool {
	return game.checkForAndRemoveCapturedPiecesOfPlayer(game.currentPlayer)
}
