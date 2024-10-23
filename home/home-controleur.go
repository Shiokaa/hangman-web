package home

import (
	"HangmanWeb/templates"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data := Home{FormDifficulty: r.FormValue("difficulty"), Difficulty: 0, Pseudo: r.FormValue("pseudo"), CheckPseudo: false}
	data.setDifficulty()

	message := r.FormValue("message")

	if message == "Pseudo Invalide" {
		data.CheckPseudo = true
	}

	templates.Templates.ExecuteTemplate(w, "home", data)
}

func homeTraitement(w http.ResponseWriter, r *http.Request) {
	data := Home{Pseudo: r.FormValue("pseudo"), FormDifficulty: r.FormValue("difficulty")}

	if !data.isValidNickname() {
		http.Redirect(w, r, "/home?message=Pseudo Invalide", http.StatusMovedPermanently)
		return
	}

	if !data.isValidDifficulty() {
		http.Redirect(w, r, "/home?message=Difficulté invalide", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
