package accounts

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //this is okay because we also called database/sql
	"github.com/zanefinner-projects/blog/config"
)

//Signup resolves /accounts/signup
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/signup.html", 301)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		fmt.Println(r.Form)
		success, message := validate(r.PostFormValue("email"), r.PostFormValue("name"), r.PostFormValue("password"), r.PostFormValue("password-repeat"))
		if success {
			fmt.Println("Signup creds are valid")
			name, pword, dbname := config.GetDbCreds()
			db, err := sql.Open("mysql", name+":"+pword+"@/"+dbname)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("DB CONNECTED")
			var insForm *sql.Stmt
			insForm, err = db.Prepare("INSERT INTO accounts(name, email, password) VALUES(?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			var res sql.Result
			res, err = insForm.Exec(r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password"))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)

		} else {
			fmt.Println("There was an invalid credential...")
			fmt.Println(message)
		}
		//Take in data
		//Validate (prevent sql injection)
		//Match to see if exists
		//If all is good, start a session
	}
	fmt.Println("ANY: /accounts/signup")
}

func validate(email, name, pword, pwordRepeat string) (bool, string) {
	if len(email) <= 14 {
		return false, "Email length is under 14 characters"
	}
	if len(name) <= 8 {
		return false, "Name is under 8 characters"
	}
	if len(pword) <= 8 {
		return false, "Password is under 8 characters"
	}
	if pword != pwordRepeat {
		return false, "Passwords did not match!"
	}
	//More validation to come!
	return true, ""
}
