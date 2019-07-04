package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckForAndRemoveCapturedPieces(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.Board[0][0] = uint8(game.CurrentPlayer.opposingPlayer())
	got := game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.Board[0][1] = uint8(game.CurrentPlayer)
	got = game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.Board[1][0] = uint8(game.CurrentPlayer) // the current player surrounds opposing piece
	got = game.checkForAndRemoveCapturedPieces()
	assert.Equal(true, got)

	assert.Equal(uint8(0), game.Board[0][0]) // the piece is removed
}

func TestCheckForAndRemoveSelfCapturedPieces(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.Board[0][1] = uint8(game.CurrentPlayer.opposingPlayer())
	got := game.checkForAndRemoveSelfCapturedPieces()
	assert.Equal(false, got)

	game.Board[1][0] = uint8(game.CurrentPlayer.opposingPlayer())
	got = game.checkForAndRemoveSelfCapturedPieces()
	assert.Equal(false, got)

	game.Board[0][0] = uint8(game.CurrentPlayer)
	got = game.checkForAndRemoveSelfCapturedPieces() // the current player plays in an invalid square
	assert.Equal(true, got)

	assert.Equal(uint8(0), game.Board[0][0]) // the piece is removed
}
