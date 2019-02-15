package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(customFileSystem{http.Dir("static")})
	mux.Handle("/", fileServer)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// don't show directory listings
// attempt to show index.html otherwise return error -> 404
// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
type customFileSystem struct {
	fs http.FileSystem
}

func (cfs customFileSystem) Open(path string) (http.File, error) {
	f, err := cfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := cfs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
