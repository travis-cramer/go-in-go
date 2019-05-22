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
