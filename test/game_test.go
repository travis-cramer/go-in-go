package main

import (
    "testing"
)

func TestChangeTurn(t *testing.T) {
    game := Game{currentMove: LIGHT}
    game.changeTurn()
    if game.currentMove != DARK {
        t.Errorf("game.currentMove = %v; want dark\n", game.currentMove)
    }
    game.changeTurn()
    if game.currentMove != LIGHT {
        t.Errorf("game.currentMove = %v; want light\n", game.currentMove)
    }
}
