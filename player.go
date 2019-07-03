package main

import "errors"
import "strings"

type Player uint8

const (
	DARK  Player = 1
	LIGHT Player = 2
)

func (player Player) toInt() uint8 {
	return uint8(player)
}

func PlayerFromInt(i uint8) (Player, error) {
	if DARK.toInt() == i {
		return DARK, nil
	}
	if LIGHT.toInt() == i {
		return LIGHT, nil
	}
	return DARK, errors.New("could not get player from int")
}

func (player Player) toString() string {
	if player == DARK {
		return "dark"
	}
	if player == LIGHT {
		return "light"
	}
	return "could not get player string"
}

func PlayerFromString(str string) (Player, error) {
	if DARK.toString() == strings.ToLower(str) {
		return DARK, nil
	}
	if LIGHT.toString() == strings.ToLower(str) {
		return LIGHT, nil
	}
	return DARK, errors.New("could not get player from string")
}

func (player Player) opposingPlayer() Player {
	if player == DARK {
		return LIGHT
	}
	if player == LIGHT {
		return DARK
	}
	return DARK
}
