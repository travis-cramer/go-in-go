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
    game.addBoardStars()
    game.printGame()

    gameOn := 1
    for gameOn == 1 {
        piecePlaced := false
        for piecePlaced == false {
            row := 0
            for row == 0 {
                row = promptForInt("Enter row: ")
            }
            col := 0
            for col == 0 {
                col = promptForInt("Enter column: ")
            }
            piecePlaced = game.placePiece(row, col, game.currentPlayer)
        }
        game.changeTurn()
        game.printGame()
    }
}

func promptForBoardSize() int {
    boardSizeAsInt := promptForInt("Board size is nxn, and must be odd. Enter n: ")
    findOdd := oddTester(int(boardSizeAsInt))
    return findOdd
}

func promptForInt(prompt string) int {
    fmt.Printf(prompt)
    var inputtedInt string
    _, err := fmt.Scanf("%v", &inputtedInt)
    if err != nil {
        log.Fatal(err)
    }
    inputtedIntAsInt, err := strconv.ParseInt(inputtedInt, 10, 8)
    if err != nil {
        fmt.Println("Error parsing input. Enter a valid integer.")
        return int(0)
    }
    return int(inputtedIntAsInt)
}

func oddTester(testInput int) int {
    isOdd := testInput % 2
    if isOdd == 1 {
        return int(testInput)
    } else {
        fmt.Println("Error, the board must have an odd number of rows and columns.")
        return int(0)
    }
}
