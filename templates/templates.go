package templates

import "html/template"

type page struct {
	Title   string
	Content string
}

//Build creates views base on template files
func Build(title, content string) bool {
	execPage := page{
		Title:   title,
		Content: content,
	}
	temp, err := template.ParseFiles('templates/index.html')
	if err != nil {
		fmt.Println(err)
		return false
	}
	temp.Execute(w, execPage)
}
