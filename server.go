package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var games []Game

func HandleRequests() {
	setUpTestData()

	http.HandleFunc("/", getGame)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func getGame(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(games[0].toGameView())
	fmt.Println("request received")
}

func setUpTestData() {
	game := NewGame(3)
	game.placePiece(1, 1, game.CurrentPlayer)
	game.placePiece(3, 3, game.CurrentPlayer.opposingPlayer())
	games = append(games, game)
}
