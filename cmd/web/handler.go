package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{ //Slice of files to parse. The "base" file should be the first
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	//URL.Query().Get() grabs the value of the id parameter from the URL query string,
	//passes it to strconv.Atoi() that converts the string to an integer value.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	//If id is not a positive integer or the id was initially
	//non-convertible to an integer, display the error message
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		//Add a "Allow" header in the response
		//header map with possible request methods
		w.Header().Set("Allow", http.MethodPost)

		//Shortcut function for displaying error message
		//by calling WriteHeader & Write BTS
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) //405 code replaced by the constant

		return
	}

	w.Write([]byte("Create a new snippet..."))
}
