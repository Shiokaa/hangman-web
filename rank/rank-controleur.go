package rank

import (
	"HangmanWeb/home"
	"HangmanWeb/templates"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func rank(w http.ResponseWriter, r *http.Request) {

	if home.Pseudo == "" {
		http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
		return
	}

	// Lire le contenu du fichier existant
	file, _ := os.ReadFile("./rank/rank.txt")

	// Écrire les nouvelles données dans le fichier
	os.WriteFile("./rank/rank.txt", []byte(string(file)+"Joueur : "+home.Pseudo+" Difficulte : "+home.Difficulty+"\n"), 0666)

	// Lire à nouveau le fichier mis à jour
	rank, _ := os.ReadFile("./rank/rank.txt")

	// Remplacer les nouvelles lignes par <br> pour HTML et convertir en template.HTML
	rankHTML := template.HTML(strings.ReplaceAll(string(rank), "\n", "<br>"))

	// Passer le contenu formaté au template
	templates.Templates.ExecuteTemplate(w, "rank", rankHTML)
}
