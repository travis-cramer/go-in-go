package main

import (
    "fmt"
    "strconv"
)

type Player uint8
const (
    DARK Player  = 1
    LIGHT Player = 2
)

func (player Player) toString() string {
    if player == DARK { return "dark" }
    if player == LIGHT { return "light" }
    return "could not get player string"
}

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
    newGame := Game{board: board, currentPlayer: DARK, gameBoardSize: boardSize}
    return newGame
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

func (game *Game) printGame() {
    fmt.Printf("current move: %v\n", game.currentPlayer.toString())

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

func (game *Game) changeTurn() {
    if game.currentPlayer == LIGHT { game.currentPlayer = DARK }
    if game.currentPlayer == DARK { game.currentPlayer = LIGHT }
}

func (game *Game) placePiece(row int, col int, player Player) bool {
    if (row > game.gameBoardSize || col > game.gameBoardSize) {
        fmt.Printf("Error, the row and column must be less than %v to fit on the game board. Please try again. \n", game.gameBoardSize)
        return false
    }
    game.board[row - 1][col - 1] = uint8(player)
    return true
}
