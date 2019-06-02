package main

import (
    "fmt"
    "log"
    "strconv"
)

func main() {
    boardSize := 0
    for boardSize == 0 {
        boardSize = promptForBoardSize()
    }
    game := NewGame(boardSize)
    game.printGame()

    gameOn := 1
    piecePlaced := false
    for gameOn == 1 { //loop for multiple moves
        for piecePlaced == false { //loop for piece placement error
            row := 0
            for row == 0 {
                row = promptForRow()
            }
            col := 0
            for col == 0 {
                col = promptForCol()
            }
            piecePlaced = game.placePiece(row, col, game.currentMove)
        }
        game.changeTurn()
        game.printGame()
    }
}

func promptForBoardSize() int {
    fmt.Printf("Board size is nxn, and must be odd. Enter n: ")
    var inputtedBoardSize string
    var oddTester int64
    _, err := fmt.Scanf("%v", &inputtedBoardSize)
    if err != nil {
        log.Fatal(err)
    }
    boardSizeAsInt, err := strconv.ParseInt(inputtedBoardSize, 10, 8)
    if err != nil {
        fmt.Println("Error parsing input. Ensure board size is a valid integer.")
        return int(0)
    }
    oddTester = boardSizeAsInt % 2
    if oddTester == 1 {
        return int(boardSizeAsInt)
    } else {
        fmt.Println("Error, the board must have an odd number of rows and columns.")
        return int(0)
    }
}

func promptForRow() int {
    fmt.Printf("Enter row: ")
    var inputtedRow string
    _, err := fmt.Scanf("%v", &inputtedRow)
    if err != nil {
        log.Fatal(err)
    }
    rowAsInt, err := strconv.ParseInt(inputtedRow, 10, 8)
    if err != nil {
        fmt.Println("Error, please enter a valid integer.")
        return int(0)
    }
    return int(rowAsInt)
}

func promptForCol() int {
    fmt.Printf("Enter col: ")
    var inputtedCol string
    _, err := fmt.Scanf("%v", &inputtedCol)
    if err != nil {
        log.Fatal(err)
    }
    colAsInt, err := strconv.ParseInt(inputtedCol, 10, 8)
    if err != nil {
        fmt.Println("Error, please enter a valid integer.")
        return int(0)
    }
    return int(colAsInt)
}
