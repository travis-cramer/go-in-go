package main

import (
    "fmt"
    "strconv"
)

type Player string
const (
    LIGHT Player = "light"
    DARK  Player = "dark"
)

type Game struct {
    board [][]uint8
    currentMove Player
}

func NewGame(boardSize int) Game {
    board := make([][]uint8, boardSize)
    for i := range board {
        board[i] = make([]uint8, boardSize)
    }
    newGame := Game{board: board, currentMove: DARK}
    return newGame
}

func (game *Game) printGame() {
    row := 0
    fmt.Printf("current move: %v\n", game.currentMove)
    for i := range game.board {
        fmt.Printf("%v  %v\n", game.board[i], row)
        row++
    }
    fmt.Printf("\n")
    for i := 0; i < row; i++ {
        iAsString := strconv.Itoa(i)
        fmt.Printf(" %v", iAsString[len(iAsString)-1:])
    }
    fmt.Printf("\n")
}

func (game *Game) changeTurn() {
    if game.currentMove == LIGHT {
        game.currentMove = DARK
    } else if game.currentMove == DARK {
        game.currentMove = LIGHT
    }
}

func (game *Game) placePiece(row int, col int, player Player) {
    if player == LIGHT {
        game.board[row][col] = 1
    } else if player == DARK {
        game.board[row][col] = 2
    }
}
