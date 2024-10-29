package home

import (
	"HangmanWeb/random"
	"math/rand/v2"
	"regexp"
)

type Home struct {
	FormDifficulty  string
	Pseudo          string
	CheckPseudo     bool
	CheckDifficulty bool
}

func (h *Home) isValidNickname() bool {
	match, _ := regexp.MatchString("^[a-zA-Z]{1,16}$", h.Pseudo)
	return match
}

func (h *Home) isValidDifficulty() bool {
	if h.FormDifficulty == "normal" || h.FormDifficulty == "difficile" {
		return true
	}
	return false
}

func CreateSlice(difficulty string) []rune {
	var runeSlice []rune
	n := rand.IntN(len(random.Word) - 1)
	x := rand.IntN(len(random.Word) - 1)

	for x == n {
		x = rand.IntN(len(random.Word) - 1)
	}

	if difficulty == "difficile" {
		for i, char := range random.Word {
			if i == len(random.Word)-1 {
				break
			}
			if i == n {
				runeSlice = append(runeSlice, char)
			} else {
				runeSlice = append(runeSlice, '_')
			}
		}
	}

	if difficulty == "normal" {
		for i, char := range random.Word {
			if i == len(random.Word)-1 {
				break
			} else if i == n {
				runeSlice = append(runeSlice, char)
			} else if i == x {
				runeSlice = append(runeSlice, char)
			} else {
				runeSlice = append(runeSlice, '_')
			}
		}
	}

	return runeSlice
}
