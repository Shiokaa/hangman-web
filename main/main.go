package main

import (
	errors "HangmanWeb/error"
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
	errors.ErrorRouteur()

	http.ListenAndServe("localhost:8080", nil)
}
