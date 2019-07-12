package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var games []Game

type Chat struct {
	Messages []Message
}

type Message struct {
	Content string
	Player  string
}

func HandleRequests() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.HandleFunc("/games", getGame)
	http.HandleFunc("/new_game", newGame)
	http.HandleFunc("/place_piece", placePiece)
	http.HandleFunc("/reset_game", resetGame)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func newGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("request received")

	boardSize := getIntegerQueryParam("board_size", r, w)

	newGame := NewGame(boardSize)
	newGame.gameOn = true
	newGame.chat = Chat{Messages: []Message{Message{Content: "hello", Player: "light"}}}
	games = append(games, newGame)

	json.NewEncoder(w).Encode(games[0].toGameView())
}

func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("request received")

	gameIdInt := getIntegerQueryParam("game_id", r, w)
	if gameIdInt >= len(games) || gameIdInt < 0 {
		fmt.Fprintf(w, "HEY! THAT GAME DOES NOT EXIST!\n")
		return
	}

	json.NewEncoder(w).Encode(games[gameIdInt].toGameView())
}

func placePiece(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("request received - place piece")

	gameIdInt := getIntegerQueryParam("game_id", r, w)
	if gameIdInt >= len(games) || gameIdInt < 0 {
		fmt.Fprintf(w, "HEY! THAT GAME DOES NOT EXIST!\n")
		return
	}
	game := &games[gameIdInt]

	row := getIntegerQueryParam("row", r, w) + 1
	col := getIntegerQueryParam("col", r, w) + 1
	onBoard := game.indexIsOnBoard(row-1, col-1)
	if !onBoard {
		fmt.Fprintf(w, "HEY! THAT INDEX IS NOT ON THE GAME BOARD!\n")
		return
	}

	piecePlaced := game.placePiece(row, col, game.currentPlayer)
	fmt.Println(piecePlaced)
	if piecePlaced {
		game.changeTurn()
	}

	json.NewEncoder(w).Encode(games[gameIdInt].toGameView())
}

func resetGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("request received")

	gameIdInt := getIntegerQueryParam("game_id", r, w)
	if gameIdInt >= len(games) || gameIdInt < 0 {
		fmt.Fprintf(w, "HEY! THAT GAME DOES NOT EXIST!\n")
		return
	}
	game := &games[gameIdInt]
	game.reset()

	json.NewEncoder(w).Encode(games[gameIdInt].toGameView())
}

func getIntegerQueryParam(fieldName string, r *http.Request, w http.ResponseWriter) int {
	keys, ok := r.URL.Query()[fieldName]

	var gameId string
	if ok && len(keys) > 0 {
		gameId = keys[0]
	}

	fmt.Printf("%s: %s\n", fieldName, gameId)
	gameIdInt, err := strconv.Atoi(gameId)
	if err != nil {
		fmt.Fprintf(w, "HEY! GIVE ME AN INTEGER FOR %s!\n", fieldName)
		return -1
	}
	return gameIdInt
}
