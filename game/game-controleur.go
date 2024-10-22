package game

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func GameControleur() {
	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	temp, err := template.ParseGlob("./game/*.html")
	if err != nil {
		fmt.Println(fmt.Printf("ERREUR => %v", err.Error()))
		os.Exit(02)
	}

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		temp.ExecuteTemplate(w, "game", nil)
	})

	http.ListenAndServe("localhost:8080", nil)
}
