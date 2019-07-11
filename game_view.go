package main

// GameView : a public representation of a Game object -- used for encoding the game object into JSON
type GameView struct {
	Board         [][]int
	BoardSize     int
	CurrentPlayer string
	GameOn        bool
}

func (game *Game) toGameView() GameView {
	boardView := NewBoardView(game.boardSize)
	for i, _ := range game.board {
		for j, _ := range game.board {
			boardView[i][j] = int(game.board[i][j])
		}
	}
	return GameView{
		Board:         boardView,
		BoardSize:     game.boardSize,
		CurrentPlayer: game.currentPlayer.toString(),
		GameOn:        game.gameOn}
}

func NewBoardView(boardSize int) [][]int {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	return board
}
