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

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

	templates.Templates.ExecuteTemplate(w, "game", data)
}

func gameTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

	data.findLetterOrWord()
	data.endGame()

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameRestart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	recupDifficulty := r.FormValue("change-difficulty")

	if recupDifficulty != "Ne pas changer de difficulté" {
		home.Difficulty = recupDifficulty
	}

	Counter = 6
	Win = false
	Lose = false
	random.Word = random.RandomWord()
	home.HiddenWord = string(home.CreateSlice(home.Difficulty))
	LettersAlreadyFound = []string{}
	WordsAlreadyFound = []string{}
	AsWon = false
	BonusUsed = true

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameBonus(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

	data.useBonus()
	data.endGame()

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameChangeNickName(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	Counter = 6
	Win = false
	Lose = false
	random.Word = random.RandomWord()
	home.HiddenWord = string(home.CreateSlice(home.Difficulty))
	LettersAlreadyFound = []string{}
	WordsAlreadyFound = []string{}
	AsWon = false
	BonusUsed = true
	Score = 0

	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
