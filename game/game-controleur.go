package game

import (
	"HangmanWeb/home"
	"HangmanWeb/random"
	"HangmanWeb/templates"
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Word: random.Word}

	templates.Templates.ExecuteTemplate(w, "game", data)
}
