package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckForAndRemoveCapturedPieces(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.board[0][0] = uint8(game.currentPlayer.opposingPlayer())
	got := game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.board[0][1] = uint8(game.currentPlayer)
	got = game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.board[1][0] = uint8(game.currentPlayer) // the current player surrounds opposing piece
	got = game.checkForAndRemoveCapturedPieces()
	assert.Equal(true, got)

	assert.Equal(uint8(0), game.board[0][0]) // the piece is removed
}

func TestCheckForAndRemoveSelfCapturedPieces(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.board[0][1] = uint8(game.currentPlayer.opposingPlayer())
	got := game.checkForAndRemoveSelfCapturedPieces()
	assert.Equal(false, got)

	game.board[1][0] = uint8(game.currentPlayer.opposingPlayer())
	got = game.checkForAndRemoveSelfCapturedPieces()
	assert.Equal(false, got)

	game.board[0][0] = uint8(game.currentPlayer)
	got = game.checkForAndRemoveSelfCapturedPieces() // the current player plays in an invalid square
	assert.Equal(true, got)

	assert.Equal(uint8(0), game.board[0][0]) // the piece is removed
}
