package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil{
		fmt.Println("Failed to write response")
	} else {
		fmt.Println("hello from helloWorld endpoint!")
	}
}

func registerHandlers() {
	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", nil))
}

func main() {
	registerHandlers()
}