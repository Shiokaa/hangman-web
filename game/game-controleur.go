package game

import (
	"HangmanWeb/templates"
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {
	templates.Templates.ExecuteTemplate(w, "game", nil)
}
