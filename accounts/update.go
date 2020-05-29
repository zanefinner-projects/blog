package accounts

import (
	"fmt"
	"net/http"
)

//Update resolves /accounts/update
func Update(w http.ResponseWriter, r *http.Request) {
	//Only let this all work if a session is set
	if r.Method == http.MethodGet {

		//Show form

	} else if r.Method == http.MethodPut {

		//Take in data
		//Validate (prevent sql injection)
		//If all is good, make changes
	}
	fmt.Println("ANY: /accounts/settings")
}
