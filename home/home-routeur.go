package home

import "net/http"

func HomeRouteur() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/home/traitement", homeTraitement)
}
