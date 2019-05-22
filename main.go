package main

import (
    "fmt"
    "log"
    "strconv"
)

func main() {
    boardSize := promptForBoardSize()
    game := NewGame(boardSize)
    game.printGame()
    row := promptForRow()
    col := promptForCol()
    game.placePiece(row, col, game.currentMove)
    game.changeTurn()
    game.printGame()
}

func promptForBoardSize() int {
    fmt.Printf("Board size is nxn. Enter n: ")
    var inputtedBoardSize string
    _, err := fmt.Scanf("%v", &inputtedBoardSize)
    if err != nil {
        log.Fatal(err)
    }
    boardSizeAsInt, err := strconv.ParseInt(inputtedBoardSize, 10, 8)
    if err != nil {
        fmt.Println("Error parsing input. Ensure board size is a valid integer.")
        log.Fatal(err)
    }
    return int(boardSizeAsInt)
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
        log.Fatal(err)
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
        log.Fatal(err)
    }
    return int(colAsInt)
}
