package accounts

import (
	"fmt"
	"net/http"
)

//Login resolves /accounts/login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		//Show form

	} else if r.Method == http.MethodPost {

		//Take in data
		//Validate (prevent sql injection)
		//Match
		//If all is good, start a session
	}
	fmt.Println("ANY: /accounts/login")
}
