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

var rankCounter int = 0

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
		isChanged := false

		for i, elem := range lines {
			if strings.Contains(elem, home.Pseudo) && strings.Contains(elem, home.Difficulty) {

				parts := strings.Split(elem, ":")
				fileScoreStr := strings.TrimSpace(parts[3])
				fileScore, _ := strconv.Atoi(fileScoreStr)

				if game.Score > fileScore {
					lines[i] = "Joueur : " + home.Pseudo + ",    Difficulté : " + home.Difficulty + ",    Score : " + score
				}

				isChanged = true
				break
			}
		}

		if !isChanged {
			lines = append(lines, "Joueur : "+home.Pseudo+",    Difficulté : "+home.Difficulty+",    Score : "+score)
		}

		if rankCounter == 0 {
			newContent := strings.Join(lines, "")
			os.WriteFile("./rank/rank.txt", []byte(newContent), 0666)
			rankCounter = 1
		} else {
			newContent := strings.Join(lines, "\n")
			os.WriteFile("./rank/rank.txt", []byte(newContent), 0666)
		}
	}

	rank, _ := os.ReadFile("./rank/rank.txt")
	data = template.HTML(strings.ReplaceAll(string(rank), "\n", "<br>"))

	templates.Templates.ExecuteTemplate(w, "rank", data)
}
