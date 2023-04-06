package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	add := r.URL.Query().Get("add")
	wait := r.URL.Query().Get("wait")

	addNumber, err := strconv.Atoi(add)
	if err != nil {
		addNumber = 1
	}

	waitNumber, err := strconv.Atoi(wait)
	if err != nil {
		waitNumber = 1
	}
	var wg sync.WaitGroup
	wg.Add(addNumber)

	for i := 0; i < waitNumber; i++ {
		wg.Done()
	}
	wg.Wait()
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	log.Println("start service:8000")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
