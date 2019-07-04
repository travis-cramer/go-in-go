package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	assert := asserter.New(t)

	desiredBoardSize := 19
	game := NewGame(desiredBoardSize)

	assert.Equal(desiredBoardSize, game.BoardSize)
	assert.Equal(desiredBoardSize, len(game.Board))
	assert.Equal(desiredBoardSize, len(game.Board[0]))
}

func TestPlacePiece(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.placePiece(1, 1, game.CurrentPlayer)
	assert.Equal(uint8(game.CurrentPlayer), game.Board[0][0])
}

func TestChangeTurn(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)
	game.CurrentPlayer = DARK

	game.changeTurn()
	assert.Equal(LIGHT, game.CurrentPlayer)

	game.changeTurn()
	assert.Equal(DARK, game.CurrentPlayer)
}
