package rank

import (
	"HangmanWeb/home"
	"HangmanWeb/templates"
	"net/http"
	"os"
)

func rank(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile()
	os.WriteFile("./rank/rank.txt", []byte("Joueur : "+home.Pseudo+" Difficulté : "+home.Difficulty), 0666)

	templates.Templates.ExecuteTemplate(w, "rank", nil)
}
