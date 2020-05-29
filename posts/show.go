package posts

import (
	"fmt"
	"net/http"
)

//ShowAll resolves /posts
func ShowAll(w http.ResponseWriter, r *http.Request) {
	//Show posts newest - > oldest
	fmt.Println("ANY: /posts")
}

//Show resolves /posts/id
func Show(w http.ResponseWriter, r *http.Request, id int) {
	//show based on id
	fmt.Println("ANY: /posts/{id}")
}
