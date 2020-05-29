package main

import (
	"log"
	"net/http"

	"github.com/zanefinner-templates/blog/accounts"
)

func main() {

	//Handlers
	///Base
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/hello.html", 301)
	})
	//Static Pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//Accounts
	http.HandleFunc("/accounts/login/", accounts.Login)     //READ
	http.HandleFunc("/accounts/signup/", accounts.Signup)   //CREATE
	http.HandleFunc("/accounts/settings/", accounts.Update) //UPDATE
	http.HandleFunc("/accounts/delete/", accounts.Delete)   //DELETE
	log.Fatal(http.ListenAndServe(":3000", nil))
}
