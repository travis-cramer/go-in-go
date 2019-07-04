package main

// GameView : a public representation of a Game object -- board is of int (not uint8), current player is string (not Player nor uint8), etc.
type GameView struct {
	Board         [][]int
	BoardSize     int
	CurrentPlayer string
	GameOn        bool
}

func (game *Game) toGameView() GameView {
	boardView := NewBoardView(game.BoardSize)
	for i, _ := range game.Board {
		for j, _ := range game.Board {
			boardView[i][j] = int(game.Board[i][j])
		}
	}
	return GameView{
		Board:         boardView,
		BoardSize:     game.BoardSize,
		CurrentPlayer: game.CurrentPlayer.toString(),
		GameOn:        game.gameOn}
}

func NewBoardView(boardSize int) [][]int {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	return board
}
