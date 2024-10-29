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
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", ")}

	templates.Templates.ExecuteTemplate(w, "game", data)
}

func gameTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", ")}

	data.findLetterOrWord()
	data.endGame()

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameRestart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	home.Difficulty = r.FormValue("change-difficulty")
	Counter = 6
	Win = false
	Lose = false
	random.Word = random.RandomWord()
	home.HiddenWord = string(home.CreateSlice(home.Difficulty))
	LettersAlreadyFound = []string{}
	WordsAlreadyFound = []string{}

	if home.Difficulty != "Difficile" && home.Difficulty != "Normal" {
		http.Redirect(w, r, "/game?message=Difficulté invalide", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
