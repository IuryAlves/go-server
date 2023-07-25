package main

import (
	"fmt"
	"github.com/IuryAlves/go-server/web"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/todo", web.TODOHandler)
	http.HandleFunc("/todo/{id}", web.TODOHandler)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
