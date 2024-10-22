package home

import (
	"net/http"
	"regexp"
)

type Home struct {
	FormDifficulty string
	Difficulty     int
	Nickname       string
}

func TraitementDonnee() {
	http.HandleFunc("/home/traitement", func(w http.ResponseWriter, r *http.Request) {

		checkNickname, _ := regexp.MatchString("^[a-zA-Z]{1,16}$", r.FormValue("pseudo"))

		if !checkNickname {

		}

		http.Redirect(w, r, "/game", http.StatusSeeOther)
	})
}

func (h *Home) SetDifficulty() {
	if h.FormDifficulty == "Difficile" {
		h.Difficulty = 1
	}
}
