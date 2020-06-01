package posts

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //this is okay because we also called database/sql
	"github.com/zanefinner-templates/blog/config"
)

//Create resolves /posts/create
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("GET: /posts/create")
		http.Redirect(w, r, "/static/new-post.html", 302)

	} else if r.Method == http.MethodPost {
		fmt.Println("POST: /posts/create")
		loggedIn := matchAuth(r.PostFormValue("email"), r.PostFormValue("password"), r.PostFormValue("key"))
		if loggedIn {
			r.ParseForm()
			success, message := validate(r.PostFormValue("title"), r.PostFormValue("content"))
			if success {
				name, pword, dbname := config.GetDbCreds()
				db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
				if err != nil {
					fmt.Println(err)
					return
				}
				var insForm *sql.Stmt
				var res sql.Result
				insForm, err = db.Prepare("INSERT INTO posts(title, content) VALUES(?,?)")
				res, err = insForm.Exec(r.PostFormValue("title"), r.PostFormValue("content"))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(res)

			} else {
				fmt.Println(message)
			}
		}
	}
}

func validate(title, content string) (bool, string) {
	if len(title) < 8 || len(title) >= 127 || len(content) < 16 || len(content) > 3400 {
		return false, "Title must be 8-128 chars, Content must be 17-3400 chars"
	}
	return true, ""
}

func matchAuth(email, password, key string) bool {
	id := 0
	name, pword, dbname := config.GetDbCreds()
	db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if key == config.GetPostKey() {
		err = db.QueryRow("SELECT id FROM accounts where email = ? and password = ?", email, password).Scan(&id)
		if id != 0 {
			return true
		}
	}
	return false
}
