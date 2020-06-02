package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	Title   string
	Content template.HTML
}

//Build creates views base on template files
func Build(w http.ResponseWriter, r *http.Request, title string, content template.HTML) bool {
	execPage := page{
		Title:   title,
		Content: content,
	}
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		return false
	}
	temp.Execute(w, execPage)
	return true
}
