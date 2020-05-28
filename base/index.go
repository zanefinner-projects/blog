package base

import (
	"fmt"
	"net/http"
)

//Index handles the landing page
func Index(w http.ResponseWriter, r *http.Request) {
	go fmt.Println("ANY: index")
	fmt.Fprintf(w, "Hello!")
}
