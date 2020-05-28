package main

import (
	"log"
	"net/http"

	"github.com/zanefinner-templates/blog/base"
)

func main() {
	//Handlers
	http.HandleFunc("/", base.Index)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
