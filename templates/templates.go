package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Templates *template.Template

func Parse() {
	var err error
	Templates, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Printf("ERREUR => %v", err.Error()))
		os.Exit(02)
	}
}
