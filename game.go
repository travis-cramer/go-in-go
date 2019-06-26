package main

type Game struct {
    board [][]uint8
    currentPlayer Player
    gameBoardSize int
    gameOn bool
}

func NewGame(boardSize int) Game {
    board := NewBoard(boardSize)
    return Game{board: board, currentPlayer: DARK, gameBoardSize: boardSize}
}

func NewBoard(boardSize int) [][]uint8 {
    board := make([][]uint8, boardSize)
    for i := range board {
        board[i] = make([]uint8, boardSize)
    }
    return board
}
