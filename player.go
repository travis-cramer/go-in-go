package main

type Player uint8

const (
	DARK  Player = 1
	LIGHT Player = 2
)

func (player Player) toString() string {
	if player == DARK {
		return "dark"
	}
	if player == LIGHT {
		return "light"
	}
	return "could not get player string"
}
