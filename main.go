package main

import (
	"log"
	"net/http"
	"todo-application/endpoint"
)

func main() {
	router := endpoint.Routes()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}
}
