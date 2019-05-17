package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

//StartServer function HTTP
func StartServer() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServ: ", err)
	}
	fmt.Println("Start http at 9000")
}
