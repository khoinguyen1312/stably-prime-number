package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"stably.khoinguyen1312.com/prime/sieve"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func findLowerPrimeNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input := vars["input"]

	inputNumber, error := strconv.Atoi(input)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	result, _ := sieve.FindSmallerPrimeNumber(inputNumber)

	response := make(map[string]int)
	response["result"] = result

	jsonResponse, error := json.Marshal(response)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./dist"))

	r.PathPrefix("/resources/").Handler(fs)
	r.HandleFunc("/api/findLowerPrimeNumber/{input}", findLowerPrimeNumber).Methods("GET")
	r.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT") // Heroku provides the port to bind to
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
