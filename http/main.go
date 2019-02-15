package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
