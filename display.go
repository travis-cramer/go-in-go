package main

import (
	"fmt"
	"strconv"
)


func (game *Game) printGame() {
	fmt.Printf("\n" + "current move: %v\n", game.currentPlayer.toString())

	row := 1
	for i := range game.board {
		fmt.Printf("%v  %v\n", game.board[i], row)
		row++
	}
	fmt.Printf("\n")

	for i := 1; i < row; i++ {
		iAsString := strconv.Itoa(i)
		fmt.Printf(" %v", iAsString[len(iAsString)-1:])
	}
	fmt.Printf("\n")
}

func (game *Game) addBoardStars() {
	center := game.gameBoardSize / 2
	if game.gameBoardSize < 12 {
		game.board [center][center] = 8
	} else if (game.gameBoardSize > 12 && game.gameBoardSize < 18) {
		game.board [3][3] = 8
		game.board [3][game.gameBoardSize - 4] = 8
		game.board [game.gameBoardSize - 4][3] = 8
		game.board [game.gameBoardSize - 4][game.gameBoardSize - 4] = 8
		game.board [center][center] = 8
	} else if (game.gameBoardSize > 18) {
		game.board [3][3] = 8
		game.board [3][center] = 8
		game.board [3][game.gameBoardSize - 4] = 8
		game.board [center][3] = 8
		game.board [center][center] = 8
		game.board [center][game.gameBoardSize - 4] = 8
		game.board [game.gameBoardSize - 4][3] = 8
		game.board [game.gameBoardSize - 4][center] = 8
		game.board [game.gameBoardSize - 4][game.gameBoardSize - 4] = 8
	}
}
