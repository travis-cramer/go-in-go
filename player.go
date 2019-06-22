package main

type Player uint8
const (
	DARK  Player = 1
	LIGHT Player = 2
)

func PlayerFromInt(int uint8) Player {
	if uint8(LIGHT) == int {
		return LIGHT
	}
	return DARK
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
