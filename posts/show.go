package posts

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/zanefinner-templates/blog/config"
)

//PostsInList struct for list view items
type postsInList struct {
	ID    int
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
		fmt.Fprintf(w, title, content)
	} else {

		name, pword, dbname := config.GetDbCreds()
		db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		results, err := db.Query("SELECT id, title FROM posts")
		if err != nil {
			fmt.Println(err)
		}

		for results.Next() {
			var post postsInList
			// for each row, scan the result into our tag composite object
			err = results.Scan(&post.ID, &post.Title)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(w, post.Title, "<hr>")
		}
	}
}
