package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	desiredBoardSize := 19
	game := NewGame(desiredBoardSize)
	if game.gameBoardSize != desiredBoardSize {
		t.Errorf("game.gameBoardSize = %d; want %d", game.gameBoardSize, desiredBoardSize)
	}
}

func TestPlacePiece(t *testing.T) {
	game := NewGame(19)
	game.placePiece(1, 1, game.currentPlayer)
	if game.board[0][0] != uint8(game.currentPlayer) {
		t.Errorf("piece on row 1 col 1 = %d; want %d", game.board[0][0], game.currentPlayer)
	}
}

func TestChangeTurn(t *testing.T) {
	game := Game{currentPlayer: LIGHT}
	game.changeTurn()
	if game.currentPlayer != DARK {
		t.Errorf("game.currentPlayer = %v; want %v\n", game.currentPlayer.toString(), DARK.toString())
	}
	game.changeTurn()
	if game.currentPlayer != LIGHT {
		t.Errorf("game.currentPlayer = %v; want %v\n", game.currentPlayer.toString(), LIGHT.toString())
	}
}
