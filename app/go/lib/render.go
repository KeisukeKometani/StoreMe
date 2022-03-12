package lib

import (
	"net/http"
	"html/template"

)

func RenderCommon(w http.ResponseWriter, r *http.Request, data interface{}, fileNames ...string) {
	t, err := template.ParseFiles(fileNames...)
	ErrorCheck(err)

	t.Execute(w, data)
}