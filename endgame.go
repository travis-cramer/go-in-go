package main

import (
	"fmt"
	"strconv"
)

type Area struct {
	darkArea          bool
	lightArea         bool
}

type Territory uint8

const (
	DARKTERRITORY    Territory = 7
	LIGHTTERRITORY   Territory = 4
	NEUTRALTERRITORY Territory = 5
)

func (territory Territory) toInt() uint8 {
	return uint8(territory)
}

func newArea() Area {
	area := Area{darkArea: false, lightArea: false}
	return area
}

func (game *Game) gameOver(hiddenGameBoard [][] uint8) {
	hiddenGameBoard = game.territoryBuilder(hiddenGameBoard)
	game.board = hiddenGameBoard
	p1, p2 := game.gameTally()
	game.printFinalBoard(p1, p2)
}

func (game *Game) territoryBuilder(hiddenGameBoard [][] uint8) [][]uint8 { //find an unidentified territory, and begin constructing a new area
	for i, _ := range game.board {
		for j, _ := range game.board {
			if (hiddenGameBoard[i][j] == 0 || hiddenGameBoard[i][j] == 8) {
				area := newArea()
				hiddenGameBoard = area.territoryBuilderHelper(hiddenGameBoard, i, j, game)
				hiddenGameBoard = area.territoryLabeler(hiddenGameBoard, i, j, game)
			}
		}
	}
	return hiddenGameBoard
}

func (area *Area) territoryBuilderHelper (hiddenGameBoard [][] uint8, index1 int, index2 int, game *Game) [][]uint8 { //Identify area with a temporary label, and find which player borders it
	indexesOfAdjacentSpaces := [4][2]int{{index1 - 1, index2}, {index1 + 1, index2}, {index1, index2 - 1}, {index1, index2 + 1}}
	hiddenGameBoard [index1][index2] = 9 
	for _, index := range indexesOfAdjacentSpaces {
		if game.indexIsOnBoard(index[0], index[1]) {
			i := hiddenGameBoard [index[0]][index[1]]
			switch i {
			case 0:
				hiddenGameBoard = area.territoryBuilderHelper(hiddenGameBoard, index[0], index[1], game)
			case 1:
				area.darkArea = true
			case 2:
				area.lightArea = true
			case 8:
				hiddenGameBoard = area.territoryBuilderHelper(hiddenGameBoard, index[0], index[1], game)
			}			
		}
	}
	return hiddenGameBoard
}

func (area *Area) territoryLabeler (hiddenGameBoard [][] uint8, index1 int, index2 int, game *Game) [][]uint8 { //using the border identifiers, relabel the premade territory as DARK, LIGHT, or neutral
	indexesOfAdjacentSpaces := [4][2]int{{index1 - 1, index2}, {index1 + 1, index2}, {index1, index2 - 1}, {index1, index2 + 1}}
	if (area.darkArea && !area.lightArea) {
		hiddenGameBoard [index1][index2] = DARKTERRITORY.toInt()
	} else if (area.lightArea && !area.darkArea) {
		hiddenGameBoard [index1][index2] = LIGHTTERRITORY.toInt()
	} else if (area.darkArea && area.lightArea) {
		hiddenGameBoard [index1][index2] = NEUTRALTERRITORY.toInt()
	} else if (!area.darkArea && !area.lightArea) {
		hiddenGameBoard [index1][index2] = NEUTRALTERRITORY.toInt()
		fmt.Printf("Error, a territory was not bordered by any pieces, and could not be tallied. \n")
	}
	for _, index := range indexesOfAdjacentSpaces {
		if game.indexIsOnBoard(index[0], index[1]) {
			if (hiddenGameBoard [index[0]][index[1]] == 9) {
				hiddenGameBoard = area.territoryLabeler(hiddenGameBoard, index[0], index[1], game)
			}		
		}
	}
	return hiddenGameBoard
}

func (game *Game) gameTally() (int, int) { 
	p1Score := game.p2PiecesLost
	p2Score := game.p1PiecesLost
	for i, _ := range game.board {
		for j, _ := range game.board {
			if (game.board[i][j] == DARKTERRITORY.toInt() || game.board[i][j] == DARK.toInt()) {
				p1Score += 1
			} else if (game.board[i][j] == LIGHTTERRITORY.toInt() || game.board[i][j] == LIGHT.toInt()) {
				p2Score += 1
			}
		}
	}
	return p1Score, p2Score
}

func (game *Game) printFinalBoard(p1Final int, p2Final int) {
	fmt.Printf("\n"+"Final Tally:   Player 1 (dark): %v, Player 2 (light): %v\n The final territories are labeled as follows: \n", strconv.Itoa(p1Final), strconv.Itoa(p2Final))

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
