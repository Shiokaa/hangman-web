package rank

import (
	"HangmanWeb/home"
	"HangmanWeb/templates"
	"net/http"
	"os"
)

func rank(w http.ResponseWriter, r *http.Request) {
	file, _ := os.ReadFile("./rank/rank.txt")
	os.WriteFile("./rank/rank.txt", []byte(string(file)+"Joueur : "+home.Pseudo+" Difficulté : "+home.Difficulty+"\n"), 0666)

	templates.Templates.ExecuteTemplate(w, "rank", nil)
}
