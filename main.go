package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zanefinner-projects/blog/accounts"
	"github.com/zanefinner-projects/blog/posts"
	"github.com/zanefinner-projects/blog/templates"
)

func main() {
	//Handlers
	///Base
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fmt.Fprint(w, "Not found :/")
			return
		}
		templates.Build(w, r, "Hello", "Content")
	})

	//Static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Accounts
	http.HandleFunc("/accounts/login/", accounts.Login)     //READ
	http.HandleFunc("/accounts/signup/", accounts.Signup)   //CREATE
	http.HandleFunc("/accounts/settings/", accounts.Update) //UPDATE*
	http.HandleFunc("/accounts/delete/", accounts.Delete)   //DELETE*

	//Posts
	http.HandleFunc("/posts/create/", posts.Create)    //CREATE
	http.HandleFunc("/posts/", posts.Show)             //READ
	http.HandleFunc("/posts/:id/delete", posts.Delete) //DELETE
	http.HandleFunc("/posts/:id/update", posts.Update) //UPDATE

	log.Fatal(http.ListenAndServe(":8000", nil))
}
