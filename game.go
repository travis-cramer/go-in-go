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

func (game *Game) changeTurn() {
    if game.currentPlayer == LIGHT { game.currentPlayer = DARK } else
    if game.currentPlayer == DARK { game.currentPlayer = LIGHT }
}

func (game *Game) placePiece(row int, col int, player Player) bool {
    if !game.indexIsOnBoard(row - 1, col - 1) {
        fmt.Printf("Error, the row and column must be positive and less than or equal to %v to fit on the game board. Please try again.\n", game.gameBoardSize)
        return false
    }
    game.board[row - 1][col - 1] = uint8(player)
    return true
}

func (game *Game) checkForCapturedPieces() bool {
    opposingPlayer := game.currentPlayer.opposingPlayer()
    return game.checkForCapturedPiecesOfPlayer(opposingPlayer)
}

func (game *Game) checkForCapturedPiecesOfPlayer(player Player) bool {
    for i := 0; i < game.gameBoardSize; i++ {
        for j := 0; i < game.gameBoardSize; j++ {
            if game.board[i][j] == uint8(player) {
                if game.checkIfHasLiberty(i, j, NewBoard(game.gameBoardSize)) {
                    return true
                }
            }
        }
    }
    return false
}

func (game *Game) indexIsOnBoard(i int, j int) bool {
    if i >= game.gameBoardSize || i < 0 {
        return false
    }
    if j >= game.gameBoardSize || j < 0 {
        return false
    }
    return true
}

func (game *Game) checkIfHasLiberty(i int, j int, alreadyVisited [][]uint8) bool {
    if !game.indexIsOnBoard(i, j) {
        return false
    }
    if alreadyVisited[i][j] == 1 {
        return false
    }
    alreadyVisited[i][j] = 1

    indexesOfAdjacentSpaces := [4][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}

    // if this piece is touching an open point, then return true
    for _, index := range indexesOfAdjacentSpaces {
        if game.indexIsOnBoard(index[0], index[1]) && game.board[index[0]][index[1]] == 0 {
            return true
        }
    }
    // if this piece is connected to a piece that is touching an open point, then return true
    for _, index := range indexesOfAdjacentSpaces {
        if game.checkIfHasLiberty(index[0], index[1], alreadyVisited) {
            return true
        }
    }
    return false
}
