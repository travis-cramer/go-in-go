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
