package rank

import (
	"HangmanWeb/game"
	"HangmanWeb/home"
	"HangmanWeb/templates"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func rank(w http.ResponseWriter, r *http.Request) {

	var data template.HTML

	if home.Pseudo == "" {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusSeeOther)
		return
	}

	if game.AsWon {
		file, _ := os.ReadFile("./rank/rank.txt")
		fileContent := string(file)
		lines := strings.Split(fileContent, "\n")
		score := strconv.Itoa(game.Score)

		for _, elem := range lines {
			if strings.Contains(elem, home.Pseudo) && strings.Contains(elem, home.Difficulty) {

			} else {
				os.WriteFile("./rank/rank.txt", []byte(string(file)+"Joueur : "+home.Pseudo+",    Difficulté : "+home.Difficulty+",    Score : "+score+"\n"), 0666)

				rank, _ := os.ReadFile("./rank/rank.txt")

				data = template.HTML(strings.ReplaceAll(string(rank), "\n", "<br>"))
			}
		}

	} else {
		rank, _ := os.ReadFile("./rank/rank.txt")
		data = template.HTML(strings.ReplaceAll(string(rank), "\n", "<br>"))
	}

	templates.Templates.ExecuteTemplate(w, "rank", data)
}
