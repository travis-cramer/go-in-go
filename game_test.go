package main

import (
	asserter "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	assert := asserter.New(t)

	desiredBoardSize := 19
	game := NewGame(desiredBoardSize)

	assert.Equal(desiredBoardSize, game.gameBoardSize)
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

func TestCheckForAndRemoveCapturedPieces(t *testing.T) {
	assert := asserter.New(t)

	game := NewGame(19)

	game.board[0][0] = uint8(game.currentPlayer.opposingPlayer())
	got := game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.board[0][1] = uint8(game.currentPlayer)
	got = game.checkForAndRemoveCapturedPieces()
	assert.Equal(false, got)

	game.board[1][0] = uint8(game.currentPlayer)
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
	got = game.checkForAndRemoveSelfCapturedPieces()
	assert.Equal(true, got)

	assert.Equal(uint8(0), game.board[0][0]) // the piece is removed
}
