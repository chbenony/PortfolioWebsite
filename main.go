package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the homepage!")
	http.ServeFile(w, r, r.URL.Path[1:])

	fmt.Println("Endpoint Hit: homepage")
}


func main() {
	//handleRequests()
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":10100", nil))
}