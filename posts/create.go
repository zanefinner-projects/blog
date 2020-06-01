package posts

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //this is okay because we also called database/sql
)

//Create resolves /posts/create
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/new-post.html", 302)

	} else if r.Method == http.MethodPost {
		fmt.Println("post req")
	}
	fmt.Println("ANY: /posts/create")
}
