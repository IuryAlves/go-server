package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/IuryAlves/go-server/api"
)

func main() {

	http.HandleFunc("/hello", getHandler)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}