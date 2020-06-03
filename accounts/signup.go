package accounts

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql" //this is okay because we also called database/sql
	"github.com/zanefinner-projects/blog/config"
	"github.com/zanefinner-projects/blog/templates"
)

//Signup resolves /accounts/signup
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		form := `
			<form method="post" action="/accounts/signup/">
			
			<input type="text" name="email" placeholder="Email"/><br>
			<input type="text" name="name" placeholder="Full Name"/><br>
			<input type="password" name="password" placeholder="Password"/><br>
			<input type="password" name="password-repeat" placeholder="Repeat Password"/><br>
			<button class="btn btn-secondary" type="submit">Sign Up</button>
			</form>`

		templates.Build(w, r, "Sign Up", template.HTML(form))

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

			_, err = insForm.Exec(r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password"))
			if err != nil {
				fmt.Println(err)
				return
			}
			templates.Build(w, r, "Signup Successful", template.HTML("You can now post! <br/><a href='/'>home</a>"))

		} else {
			templates.Build(w, r, "Signup Error", template.HTML(message+"<br/><a href='/accounts/signup/'>try again</a>"))
		}
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
