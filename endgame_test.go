package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestGameOver(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)
	game.board[1][0] = uint8(game.currentPlayer)
	game.board[0][1] = uint8(game.currentPlayer)
	game.board [17][18] = uint8(game.currentPlayer.opposingPlayer())
	game.board [18][17] = uint8(game.currentPlayer.opposingPlayer())
	
	game.gameOver(game.board)
	p1, p2 := game.gameTally()
	assert.Equal(3, p1) //test tallies
	assert.Equal(3, p2)
	assert.Equal(uint8(5), uint8(game.board[10][10])) //test neutral territory
}