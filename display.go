package main

import (
	"fmt"
	"strconv"
)

func (game *Game) printGame() {
	fmt.Printf("\n"+"current move: %v\n", game.CurrentPlayer.toString())

	row := 1
	for i := range game.Board {
		fmt.Printf("%v  %v\n", game.Board[i], row)
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
	center := game.BoardSize / 2
	if game.BoardSize < 12 {
		game.Board[center][center] = 8
	} else if game.BoardSize > 12 && game.BoardSize < 18 {
		game.Board[3][3] = 8
		game.Board[3][game.BoardSize-4] = 8
		game.Board[game.BoardSize-4][3] = 8
		game.Board[game.BoardSize-4][game.BoardSize-4] = 8
		game.Board[center][center] = 8
	} else if game.BoardSize > 18 {
		game.Board[3][3] = 8
		game.Board[3][center] = 8
		game.Board[3][game.BoardSize-4] = 8
		game.Board[center][3] = 8
		game.Board[center][center] = 8
		game.Board[center][game.BoardSize-4] = 8
		game.Board[game.BoardSize-4][3] = 8
		game.Board[game.BoardSize-4][center] = 8
		game.Board[game.BoardSize-4][game.BoardSize-4] = 8
	}
}
