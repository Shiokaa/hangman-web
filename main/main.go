package main

import (
	"HangmanWeb/game"
	"HangmanWeb/home"
	"HangmanWeb/templates"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	templates.Parse()

	home.HomeRouteur()
	game.GameRouteur()

	http.ListenAndServe("localhost:8080", nil)
}
