package main

import "net/http"

func main() {
	//Handlers
	http.HandleFunc("/", base.Index)
}
