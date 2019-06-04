package main

import (
    "fmt"
)

type Game struct {
    board [][]uint8
    currentPlayer Player
    gameBoardSize int
}

func NewGame(boardSize int) Game {
    board := make([][]uint8, boardSize)
    for i := range board {
        board[i] = make([]uint8, boardSize)
    }
    return Game{board: board, currentPlayer: DARK, gameBoardSize: boardSize}
}

func (game *Game) changeTurn() {
    if game.currentPlayer == LIGHT { game.currentPlayer = DARK }
    if game.currentPlayer == DARK { game.currentPlayer = LIGHT }
}

func (game *Game) placePiece(row int, col int, player Player) bool {
    if (row > game.gameBoardSize || col > game.gameBoardSize) {
        fmt.Printf("Error, the row and column must be less than or equal to %v to fit on the game board. Please try again.\n", game.gameBoardSize)
        return false
    }
    game.board[row - 1][col - 1] = uint8(player)
    return true
}
