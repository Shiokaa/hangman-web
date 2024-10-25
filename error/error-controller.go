package errors

import (
	"fmt"
	"net/http"
)

func error(w http.ResponseWriter, r *http.Request) {
	code, message := r.FormValue("code"), r.FormValue("message")
	if code != "" && message != "" {
		fmt.Fprintf(w, "Erreur %s - %s", code, message)
		return
	}
	fmt.Fprint(w, "Oups une erreur serveur est survenue")
}
