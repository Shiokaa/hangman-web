package home

import "net/http"

type Home struct {
	FormDifficulty string
	Difficulty     int
	Nickname       string
}

func TraitementDonnee() {
	http.HandleFunc("/home/traitement", func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, "/game", http.StatusSeeOther)
	})
}

func (h *Home) SetDifficulty() {
	if h.FormDifficulty == "Difficile" {
		h.Difficulty = 1
	}
}
