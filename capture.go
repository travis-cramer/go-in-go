package main

import "log"

func (game *Game) checkForAndRemoveCapturedPieces() bool {
	opposingPlayer := game.currentPlayer.opposingPlayer()
	return game.checkForAndRemoveCapturedPiecesOfPlayer(opposingPlayer)
}

func (game *Game) checkForAndRemoveSelfCapturedPieces() bool {
	return game.checkForAndRemoveCapturedPiecesOfPlayer(game.currentPlayer)
}

func (game *Game) checkForAndRemoveCapturedPiecesOfPlayer(player Player) bool {
	var piecesWereRemoved bool
	for i := 0; i < game.boardSize; i++ {
		for j := 0; j < game.boardSize; j++ {
			if game.board[i][j] == uint8(player) {
				if !game.checkIfHasLiberty(i, j, player, NewBoard(game.boardSize)) {
					game.removePieceAndAllConnectedPieces(i, j)
					piecesWereRemoved = true
				}
			}
		}
	}
	return piecesWereRemoved
}

func (game *Game) checkIfHasLiberty(i int, j int, player Player, alreadyVisited [][]uint8) bool {
	if !game.indexIsOnBoard(i, j) {
		return false
	}
	if game.board[i][j] != uint8(player) {
		return false
	}
	if alreadyVisited[i][j] == 1 {
		return false
	}
	alreadyVisited[i][j] = 1

	indexesOfAdjacentSpaces := [4][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}

	// if this piece is touching an open point, then return true
	for _, index := range indexesOfAdjacentSpaces {
		if game.indexIsOnBoard(index[0], index[1]) && game.board[index[0]][index[1]] == 0 {
			return true
		}
	}
	// if this piece is connected to a piece that is touching an open point, then return true
	for _, index := range indexesOfAdjacentSpaces {
		if game.checkIfHasLiberty(index[0], index[1], player, alreadyVisited) {
			return true
		}
	}
	return false
}

func (game *Game) removePieceAndAllConnectedPieces(i int, j int) {
	player, err := PlayerFromInt(game.board[i][j])
	if err != nil {
		log.Fatal(err)
	}
	game.removePieceAndAllConnectedPiecesHelper(i, j, player)
}

func (game *Game) removePieceAndAllConnectedPiecesHelper(i int, j int, player Player) {
	if !game.indexIsOnBoard(i, j) {
		return
	}
	if game.board[i][j] != uint8(player) {
		return
	}
	game.board[i][j] = 0 // remove piece
	game.incrimentLostPieces(player)
	
	indexesOfAdjacentSpaces := [4][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}

	for _, index := range indexesOfAdjacentSpaces {
		game.removePieceAndAllConnectedPiecesHelper(index[0], index[1], player)
	}
}

func (game *Game) incrimentLostPieces(player Player) {
	if uint8(player) == 1 {
		game.p1PiecesLost += 1
	} else if uint8(player) == 2 {
		game.p2PiecesLost += 1
	}
}