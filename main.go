package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	cliPtr := flag.Bool("cli", true, "play on cli or start api?")
	apiPtr := flag.Bool("api", false, "play on cli or start api?")
	flag.Parse()
	if *apiPtr || !*cliPtr {
		startAPI()
	} else {
		startCLI()
	}
}

func startCLI() {
	fmt.Println("Starting cli...")
	boardSize := 0
	for boardSize == 0 {
		boardSize = promptForBoardSize()
	}
	game := NewGame(boardSize)

	game.play()
}

func startAPI() {
	fmt.Println("Starting api...")
	HandleRequests()
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
		fmt.Println("Newline Detected.")
		log.Fatal(err)
	}
	if inputtedInt == strconv.Quote("\n") {
		fmt.Println("Newline Detected.")
		return int(-1)
	}
	inputtedIntAsInt, err := strconv.ParseInt(inputtedInt, 10, 8)
	if err != nil {
		fmt.Println("Newline Detected.")
		fmt.Println("Error parsing input. Enter a valid integer.")
		return int(0)
	}
	return int(inputtedIntAsInt)
}

func oddTester(testInput int) int {
	isOdd := testInput % 2
	if isOdd == 1 {
		return int(testInput)
	}
	fmt.Println("Error, the board must have an odd number of rows and columns.")
	return int(0)
}
