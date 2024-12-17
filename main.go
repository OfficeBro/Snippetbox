package main

import (
	"log"
	"net/http"

	"github.com/OfficeBro/Snippetbox/handler"
)

func main() {
	mux := http.NewServeMux() //Create a new router

	mux.HandleFunc("/", handler.Home)                    //Register home as the handler for the "\" URL pattern
	mux.HandleFunc("/snippet/view", handler.SnippetView) //Register following handlers for their specific URL patterns
	mux.HandleFunc("/snippet/create", handler.SnippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	//Create a web server
	//First parameter is the TCP network address to listen on, second param. is our servemux

	log.Fatal(err) //Display error if needed
}
