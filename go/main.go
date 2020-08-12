package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.HandleFunc("/", indexHandler)
	http.Handle("/resources/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
