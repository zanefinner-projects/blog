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

//ShowAll resolves /posts
func ShowAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ANY: /posts")
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

		fmt.Fprintln(w, post.Title)
	}
}

//Show resolves /posts/id
func Show(w http.ResponseWriter, r *http.Request) {
	//show based on id
	fmt.Println("ANY: /posts/{id}")
}
