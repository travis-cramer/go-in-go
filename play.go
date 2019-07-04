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
			piecePlaced = game.placePiece(row, col, game.CurrentPlayer)
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

	game.Board[row-1][col-1] = player.toInt()

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
	if game.CurrentPlayer == LIGHT {
		game.CurrentPlayer = DARK
	} else if game.CurrentPlayer == DARK {
		game.CurrentPlayer = LIGHT
	}
}

func (game *Game) indexIsOnBoard(i int, j int) bool {
	if i >= game.BoardSize || i < 0 {
		return false
	}
	if j >= game.BoardSize || j < 0 {
		return false
	}
	return true
}
