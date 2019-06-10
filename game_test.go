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

func TestCheckForCapturedPieces(t *testing.T) {
    assert := asserter.New(t)

    game := NewGame(19)

    game.placePiece(1, 1, game.currentPlayer)
    game.placePiece(1, 2, game.currentPlayer.opposingPlayer())
    game.placePiece(2, 1, game.currentPlayer.opposingPlayer())

    got := game.checkForCapturedPieces()

    assert.Equal(true, got)
}