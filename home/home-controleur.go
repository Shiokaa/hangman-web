package home

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func HomeControler() {
	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	temp, err := template.ParseGlob("./home/*.html")
	if err != nil {
		fmt.Println(fmt.Printf("ERREUR => %v", err.Error()))
		os.Exit(02)
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		data := Home{FormDifficulty: r.FormValue("difficulty"), Nickname: r.FormValue("pseudo")}
		data.SetDifficulty()

		temp.ExecuteTemplate(w, "home", data)
	})

	http.ListenAndServe("localhost:8080", nil)
}
