package game

import (
	"HangmanWeb/home"
	"strings"
)

type RecupVar struct {
	CheckValueInput     bool
	LetterIsRight       bool
	LetterIsWrong       bool
	Pseudo              string
	Difficulty          string
	Counter             int
	Word                string
	HiddenWord          string
	UserValue           string
	LettersAlreadyFound string
	WordsAlreadyFound   string
	Win                 bool
	Lose                bool
	Score               int
	BonusUsed           bool
}

var Counter int = 6
var LettersAlreadyFound []string
var WordsAlreadyFound []string
var Win bool
var Lose bool
var AsWon bool
var Score int
var BonusUsed bool = true
var LetterIsRight bool
var LetterIsWrong bool

func (r *RecupVar) convertedWord() []rune {
	randomWord := r.HiddenWord
	var runeSlice []rune

	for _, v := range randomWord {
		if v == 13 {
			break
		} else {
			runeSlice = append(runeSlice, v)
		}
	}

	return runeSlice
}

func (r *RecupVar) findLetterOrWord() {
	isFind := false
	hiddenWord := r.convertedWord()

	if len(r.UserValue) > 1 {
		if strings.EqualFold(strings.TrimSpace(r.UserValue), strings.TrimSpace(r.Word)) {
			r.HiddenWord = r.UserValue
			home.HiddenWord = r.UserValue
		} else {
			WordsAlreadyFound = append(WordsAlreadyFound, r.UserValue)
			Counter -= 2
		}
	} else {
		for iWord, charWord := range r.Word {
			if strings.EqualFold(string(charWord), r.UserValue) {
				isFind = true
				for iHiddenWord := range r.HiddenWord {
					if iWord == iHiddenWord {
						hiddenWord[iHiddenWord] = charWord
						home.HiddenWord = string(hiddenWord)
						r.HiddenWord = string(hiddenWord)
						LetterIsRight = true
						LetterIsWrong = false
						break
					}
				}
			}
		}

		if !isFind {
			LettersAlreadyFound = append(LettersAlreadyFound, r.UserValue)
			Counter--
			LetterIsWrong = true
			LetterIsRight = false
		}
	}

}

func (r *RecupVar) endGame() {
	if Counter == 0 {
		Lose = true
	}
	if strings.EqualFold(strings.TrimSpace(r.HiddenWord), strings.TrimSpace(r.Word)) {
		Win = true
		AsWon = true
		Score += Counter
	}
}

func (r *RecupVar) useBonus() {
	hiddenWord := r.convertedWord()
	for i, char := range r.Word {
		if hiddenWord[i] == '_' { // Choisit la première lettre non révélée
			hiddenWord[i] = char
			break
		}
	}

	r.HiddenWord = string(hiddenWord)
	home.HiddenWord = string(hiddenWord)
	BonusUsed = false
	Counter--
}
