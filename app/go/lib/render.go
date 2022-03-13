package lib

import (
	"net/http"
	"html/template"

)

func RenderCommon(w http.ResponseWriter, r *http.Request, data interface{}, fileNames ...string) {
	commonTemplateFiles := []string{"html/template/_head.html","html/template/_header.html"}
	fileNames = append(fileNames,commonTemplateFiles...)
	t, err := template.ParseFiles(fileNames...)
	ErrorCheck(err)

	t.Execute(w, data)
}