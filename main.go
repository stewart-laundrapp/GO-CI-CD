package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./static"))
		fmt.Fprintf(w, "Hello World!")
	})
}

func main() {
	// for heroku since we have to use the assigned port for the app
	port := os.Getenv("PORT")
	if port == "" {
		// if we're running on minikube or just running this with ./cicdexample
		defaultPort := "3000"
		log.Println("no env var set for port, defaulting to " + defaultPort)
		// serve the contents of the static folder on '/'
		http.Handle("/", homeHandler())
		http.ListenAndServe(":"+defaultPort, nil)
	} else {
		http.Handle("/", homeHandler())
		log.Println("starting server on port " + port)
		http.ListenAndServe(":"+port, nil)
	}
}
