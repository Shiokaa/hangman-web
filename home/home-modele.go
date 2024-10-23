package home

import (
	"regexp"
)

type Home struct {
	FormDifficulty string
	Pseudo         string
	CheckPseudo    bool
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
