package base

import (
	"fmt"
	"net/http"
)

//About returns the blog's purpose and general info
func About(w http.ResponseWriter, r *http.Request) {
	go fmt.Println("ANY: /about")
	fmt.Fprintf(w, "About us")
}
