package game

import (
	"HangmanWeb/home"
	"HangmanWeb/random"
	"HangmanWeb/templates"
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {

	if home.Pseudo == "" {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Word: random.Word, HiddenWord: home.HiddenWord, Counter: 6, UserWord: r.FormValue("word")}

	data.FindLetterOrWord()

	templates.Templates.ExecuteTemplate(w, "game", data)
}

func gameTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
		return
	}

	userword := r.FormValue("user-word")

	http.Redirect(w, r, "/game?word="+userword, http.StatusSeeOther)
}
