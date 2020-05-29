package posts

import (
	"fmt"
	"net/http"
)

//Create resolves /posts/create
func Create(w http.ResponseWriter, r *http.Request) {
	//Only let this all work if a session is set
	if r.Method == http.MethodPost { //Might have to oconvert to post

		//Validate
		//Upload
		//Redirect to article

	} else if r.Method == http.MethodGet {
		//Show creation form
	} else {
		//err
	}
	fmt.Println("ANY: /posts/create")
}
