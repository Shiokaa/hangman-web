package home

import (
	"HangmanWeb/templates"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data := Home{FormDifficulty: r.FormValue("difficulty"), Pseudo: r.FormValue("pseudo"), CheckPseudo: false}

	message := r.FormValue("message")

	if message == "Pseudo Invalide" {
		data.CheckPseudo = true
	}

	if message == "Difficulté invalide" {
		data.CheckDifficulty = true
	}

	templates.Templates.ExecuteTemplate(w, "home", data)
}

var Pseudo string
var Difficulty string
var HiddenWord string

func homeTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := Home{Pseudo: r.FormValue("pseudo"), FormDifficulty: r.FormValue("difficulty")}

	if !data.isValidNickname() {
		http.Redirect(w, r, "/home?message=Pseudo Invalide", http.StatusSeeOther)
		return
	}

	if !data.isValidDifficulty() {
		http.Redirect(w, r, "/home?message=Difficulté invalide", http.StatusSeeOther)
		return
	}

	Pseudo = r.FormValue("pseudo")
	Difficulty = data.FormDifficulty
	HiddenWord = string(CreateSlice(Difficulty))

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
