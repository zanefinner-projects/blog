package posts

import (
	"fmt"
	"net/http"
)

//Delete resolves /posts/{id}/delete
func Delete(w http.ResponseWriter, r *http.Request, int id) {
	if r.Method == http.MethodPut {
		//Validate
		//Change DB
		//Redirect to updated post
	} else if r.Method == http.MethodGet {
		//Show edit form
	}
	fmt.Println("Delete: /posts/{id}/delete")
}