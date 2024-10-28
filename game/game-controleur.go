package game

import (
	"HangmanWeb/home"
	"HangmanWeb/random"
	"HangmanWeb/templates"
	"net/http"
	"strings"
)

func game(w http.ResponseWriter, r *http.Request) {

	if home.Pseudo == "" {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", ")}

	templates.Templates.ExecuteTemplate(w, "game", data)
}

func gameTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", ")}

	data.findLetterOrWord()
	data.endGame()
	print(Win)
	print(data.Win)

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
