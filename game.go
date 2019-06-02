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
    gameBoardSize int //added board size variable for testing out of scope moves
}

func NewGame(boardSize int) Game {
    board := make([][]uint8, boardSize)
    for i := range board {
        board[i] = make([]uint8, boardSize)
    }
    board = boardStars(board, boardSize)
    newGame := Game{board: board, currentMove: DARK, gameBoardSize: boardSize}
    return newGame
}

func boardStars(board [][] uint8, boardSize int) [][] uint8 {
    center := boardSize/2
    if boardSize < 12 {
        board [center][center] = 8
    } else if (boardSize > 12 && boardSize < 18) {
        board [3][3] = 8
        board [3][boardSize - 4] = 8
        board [boardSize - 4][3] = 8
        board [boardSize - 4][boardSize - 4] = 8
        board [center][center] = 8
    } else if (boardSize > 18) {
        board [3][3] = 8
        board [3][center] = 8
        board [3][boardSize - 4] = 8
        board [center][3] = 8
        board [center][center] = 8
        board [center][boardSize - 4] = 8
        board [boardSize - 4][3] = 8
        board [boardSize - 4][center] = 8
        board [boardSize - 4][boardSize - 4] = 8
    }
    return board
}

func (game *Game) printGame() {
    row := 1 //edited to start at 1, so player labels start at 1, and so 0 can be used for errors in entering rows/columns
    fmt.Printf("current move: %v\n", game.currentMove)
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
    if game.currentMove == LIGHT {
        game.currentMove = DARK
    } else if game.currentMove == DARK {
        game.currentMove = LIGHT
    }
}

func (game *Game) placePiece(row int, col int, player Player) bool { //added int return type for testing out of bounds
    if (row > game.gameBoardSize || col > game.gameBoardSize) {
        fmt.Printf("Error, the row and column must be less than %v to fit on the game board. Please try again. \n", game.gameBoardSize)
        return false
    } else {
        if player == LIGHT {
            game.board[row][col] = 1
            return true
        } else if player == DARK {
            game.board[row][col] = 2
            return true
        }
    }
    return false
}
