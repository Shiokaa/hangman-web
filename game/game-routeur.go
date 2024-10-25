package game

import "net/http"

func GameRouteur() {
	http.HandleFunc("/game", game)
	http.HandleFunc("/game/traitement", gameTraitement)
}
