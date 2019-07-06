package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	assert := asserter.New(t)

	desiredBoardSize := 19
	game := NewGame(desiredBoardSize)

	assert.Equal(desiredBoardSize, game.boardSize)
	assert.Equal(desiredBoardSize, len(game.board))
	assert.Equal(desiredBoardSize, len(game.board[0]))
}

func TestPlacePiece(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.placePiece(1, 1, game.currentPlayer)
	assert.Equal(uint8(game.currentPlayer), game.board[0][0])
}

func TestChangeTurn(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)
	game.currentPlayer = DARK

	game.changeTurn()
	assert.Equal(LIGHT, game.currentPlayer)

	game.changeTurn()
	assert.Equal(DARK, game.currentPlayer)
}
