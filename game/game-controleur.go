package game

import (
	"HangmanWeb/home"
	"HangmanWeb/random"
	"HangmanWeb/templates"
	"net/http"
	"regexp"
	"strings"
)

func game(w http.ResponseWriter, r *http.Request) {

	if home.Pseudo == "" {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{CheckValueInput: false, LetterIsRight: LetterIsRight, LetterIsWrong: LetterIsWrong, Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

	message := r.FormValue("message")

	if message == "Input Invalide" {
		data.CheckValueInput = true
	}

	templates.Templates.ExecuteTemplate(w, "game", data)
}

func gameTraitement(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	checkValueInput, _ := regexp.MatchString("^[a-zA-Z]{1,32}$", r.FormValue("user-value"))

	if !checkValueInput {
		http.Redirect(w, r, "/game?message=Input Invalide", http.StatusMovedPermanently)
		return
	}

	data := RecupVar{CheckValueInput: false, LetterIsRight: LetterIsRight, LetterIsWrong: LetterIsWrong, Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

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
	LetterIsRight = false
	LetterIsWrong = false

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameBonus(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	data := RecupVar{CheckValueInput: false, LetterIsRight: LetterIsRight, LetterIsWrong: LetterIsWrong, Pseudo: home.Pseudo, Difficulty: home.Difficulty, Counter: Counter, Word: random.Word, HiddenWord: home.HiddenWord, UserValue: r.FormValue("user-value"), Win: Win, Lose: Lose, LettersAlreadyFound: strings.Join(LettersAlreadyFound, ", "), WordsAlreadyFound: strings.Join(WordsAlreadyFound, ", "), Score: Score, BonusUsed: BonusUsed}

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
	LetterIsRight = false
	LetterIsWrong = false

	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
