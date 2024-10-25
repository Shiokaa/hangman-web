package game

import (
	"HangmanWeb/home"
	"strings"
)

type RecupVar struct {
	Pseudo              string
	Difficulty          string
	Word                string
	HiddenWord          string
	Counter             int
	UserWord            string
	LettersAlreadyFound []string
	WordsAlreadyFound   []string
}

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

	if len(r.UserWord) > 1 {
		if strings.EqualFold(strings.TrimSpace(r.UserWord), strings.TrimSpace(r.Word)) {
			r.HiddenWord = r.UserWord
		} else {
			r.WordsAlreadyFound = append(r.WordsAlreadyFound, r.UserWord)
			r.Counter -= 2
		}
	}

	for iWord, charWord := range r.Word {
		if strings.EqualFold(string(charWord), r.UserWord) {
			isFind = true
			for iHiddenWord := range r.HiddenWord {
				if iWord == iHiddenWord {
					hiddenWord[iHiddenWord] = charWord
					home.HiddenWord = string(hiddenWord)
					r.HiddenWord = string(hiddenWord)
					break
				}
			}
		}
	}

	if !isFind {
		r.LettersAlreadyFound = append(r.LettersAlreadyFound, r.UserWord)
		r.Counter--
	}
}
