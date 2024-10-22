package game

import "net/http"

func GameRouteur() {
	http.HandleFunc("/game", game)
}
