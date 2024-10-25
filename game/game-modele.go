package game

import "HangmanWeb/random"

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

func convertWord() []rune {
	randomWord := random.Word
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

func (r *RecupVar) FindLetterOrWord() {
	/* 	isFind := false
	   	word := convertWord()
	*/
	if len(r.UserWord) > 1 {
		if r.UserWord == r.Word {
			r.HiddenWord = r.UserWord
			r.Counter += 1
		} else {
			r.WordsAlreadyFound = append(r.WordsAlreadyFound, r.UserWord)
			r.Counter -= 2
		}
	}
}
