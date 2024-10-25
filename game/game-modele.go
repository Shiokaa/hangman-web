package game

import (
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

func (r *RecupVar) FindLetterOrWord() {
	/* isFind := false */

	if len(r.UserWord) > 1 {
		if strings.EqualFold(strings.TrimSpace(r.UserWord), strings.TrimSpace(r.Word)) {
			r.HiddenWord = r.UserWord
		} else {
			r.WordsAlreadyFound = append(r.WordsAlreadyFound, r.UserWord)
			r.Counter -= 2
		}
	}
}
