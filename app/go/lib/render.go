package lib

import (
	"log"
	"net/http"
	"html/template"

)

func RenderCommon(w http.ResponseWriter, r *http.Request, data interface{}, fileNames ...string) {
	t, err := template.ParseFiles(fileNames...)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, data)
}