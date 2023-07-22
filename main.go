package main

import (
	"log"
	"net/http"

	"github.com/codeWithUtkarsh/go-abs/handler"
	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/upload", handler.Create).Methods("POST")
	myRouter.HandleFunc("/status", handler.GetStatus).Methods("GET")

	handler.SetSwagger(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
