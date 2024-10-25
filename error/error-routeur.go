package errors

import "net/http"

func ErrorRouteur() {
	http.HandleFunc("/erreur", error)
}
