package main

import (
	"log"
	"net/http"

	"github.com/zanefinner-templates/blog/base"
)

func main() {
	//Handlers
	///Base
	go http.HandleFunc("/", base.Index)
	http.HandleFunc("/about", base.About)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
