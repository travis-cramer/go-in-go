package main

// Game : representation of the current game state -- where pieces are, whose turn it is, etc.
type Game struct {
	board         [][]uint8
	boardSize     int
	currentPlayer Player
	gameOn        bool
}

// NewGame : creates a new game given a board size; chooses which player starts first and adds board stars
func NewGame(boardSize int) Game {
	board := NewBoard(boardSize)
	game := Game{board: board, currentPlayer: DARK, boardSize: boardSize}
	game.addBoardStars()
	return game
}

func NewBoard(boardSize int) [][]uint8 {
	board := make([][]uint8, boardSize)
	for i := range board {
		board[i] = make([]uint8, boardSize)
	}
	return board
}

func (game *Game) reset() {
	game.board = NewBoard(game.boardSize)
	game.currentPlayer = DARK
}
