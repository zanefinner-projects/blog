package posts

import (
	"fmt"
	"net/http"
)

//Update resolves /posts/{id}/update
func Update(w http.ResponseWriter, r *http.Request) {
	//Only let this all work if a session is set
	if r.Method == http.MethodPut { //Might have to oconvert to post

		//Update

	} else if r.Method == http.MethodGet {
		//Show update form
	} else {
		//err
	}
	fmt.Println("ANY: /posts/{id}/update")
}
