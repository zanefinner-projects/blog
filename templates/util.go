package templates

var templates *template.Template
//LoadTemplates uses a pattern to load certain file types
func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))	
}
//ExecuteTemplate calls ExecuteTemplate based on LoadTemplate's inputs
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}