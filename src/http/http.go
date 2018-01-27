package main

import (
	"net/http"
	"fmt"
)

func main() {

	//A simple http server

	http.HandleFunc("/", handleDefault)

	http.ListenAndServe(":8181", nil)
}


func handleDefault(w http.ResponseWriter, r* http.Request){
	fmt.Fprint(w, "Blubb\n")

	fmt.Println("served!")
}

