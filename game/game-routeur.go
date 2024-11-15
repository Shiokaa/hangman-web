package game

import "net/http"

func GameRouteur() {
	http.HandleFunc("/game", game)
	http.HandleFunc("/game/traitement", gameTraitement)
	http.HandleFunc("/game/restart", gameRestart)
	http.HandleFunc("/game/bonus", gameBonus)
	http.HandleFunc("/game/changeNickName", gameChangeNickName)
}
