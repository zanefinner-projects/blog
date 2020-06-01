package accounts

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/zanefinner-projects/blog/config"
)

//Login resolves /accounts/login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/login.html", 301)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		fmt.Println(r.Form)
		isLoggedIn, message, _ := match(r.PostFormValue("email"), r.PostFormValue("password"))
		if isLoggedIn {
			fmt.Println("Logged in")
			//start session
		} else {
			fmt.Println(message)
		}
	}
	fmt.Println("ANY: /accounts/login")
}
func match(email, pass string) (bool, string, int) {
	fmt.Println("Email: ", email)
	fmt.Println("Pass: ", pass)

	name, pword, dbname := config.GetDbCreds()
	db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)

	if err != nil {
		fmt.Println(err)
		return false, "DB Couldn't connect", 0
	}

	fmt.Println("DB CONNECTED")

	var uname string
	var id int

	err = db.QueryRow("SELECT id, name FROM accounts where email = ? and password = ?", email, pass).Scan(&id, &uname)

	if err != nil {
		fmt.Println(err)
		return false, "Error executing query", 0
	}

	fmt.Println(uname, id)
	return true, "", id
}
