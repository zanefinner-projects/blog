package accounts

import (
	"fmt"
	"net/http"
)

//Delete resolves /accounts/delete
func Delete(w http.ResponseWriter, r *http.Request) {
	//Only let this all work if a session is set
	if r.Method == http.MethodDelete { //Might have to oconvert to post

		//use session to "delete" account

	} else {

		//err
	}
	fmt.Println("ANY: /accounts/delete")
}
