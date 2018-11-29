package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println(params)

	name := params.Get("name")
	age := params.Get("age")

	fmt.Println(name)
	fmt.Fprintf(w, "Greetings %s", name)

	fmt.Println(age)
	if age != "" {
		fmt.Fprintf(w, " (%s)", age)
	}
}

func main() {
	http.HandleFunc("/salute", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Salute %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/greet", handleGreet)

	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
