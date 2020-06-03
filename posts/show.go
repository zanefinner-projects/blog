package posts

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/zanefinner-projects/blog/config"
	"github.com/zanefinner-projects/blog/templates"
)

//PostsInList struct for list view items
type postsInList struct {
	ID    string
	Title string
}

//Show resolves /posts
func Show(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ANY: /posts/")
	qr := r.URL.Query()
	id, ok := qr["id"]
	if ok {
		// fmt.Println("View based on id:", id[0])
		name, pword, dbname := config.GetDbCreds()
		db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
		if err != nil {
			fmt.Println(err)
			return
		}
		var title string
		var content string

		err = db.QueryRow("SELECT title, content FROM posts where id = ?", id[0]).Scan(&title, &content)
		if err != nil {
			fmt.Println(err)
		}
		templates.Build(w, r, title, template.HTML(content))
	} else {

		name, pword, dbname := config.GetDbCreds()
		db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
		var output template.HTML
		var outputString string
		results, err := db.Query("SELECT id, title FROM posts")
		if err != nil {
			fmt.Println(err)
		}
		var posts []postsInList
		itt := 0
		for results.Next() {
			itt++
			var tmp postsInList
			err = results.Scan(&tmp.ID, &tmp.Title)
			if err != nil {
				fmt.Println(err)
				return
			}
			posts = append(posts, tmp)

		}
		for key, val := range posts {
			fmt.Println(key, val)
			outputString += "<a href='/posts/?id=" + string(val.ID) + "'>" + val.Title + "</a><br>"
		}
		output = template.HTML(outputString)
		templates.Build(w, r, "Posts", output)
		fmt.Println(output)
	}
}
