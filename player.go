package main

import "errors"

type Player uint8
const (
	DARK  Player = 1
	LIGHT Player = 2
)

func PlayerFromInt(i uint8) (Player, error) {
	if uint8(DARK) == i {
		return DARK, nil
	}
	if uint8(LIGHT) == i {
		return LIGHT, nil
	}
	return DARK, errors.New("cannot get player from int")
}

func (player Player) toString() string {
	if player == DARK { return "dark" }
	if player == LIGHT { return "light" }
	return "could not get player string"
}

func (player Player) opposingPlayer() Player {
	if player == DARK { return LIGHT }
	if player == LIGHT { return DARK }
	return DARK
}
