/*
 * main.go
 */
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	//Create the data root folder
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		err = os.MkdirAll("data", 0755)
		if err != nil {
			panic(err)
		}
	}

	//Initialize API router here
	router := ChiRouter().InitRouter()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s 	%s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	//Start the server
	///TODO: start the server over https
	http.ListenAndServe(":9090", router)
	//log.Fatal( http.ListenAndServeTLS(":9090", "localhost.crt", "localhost.key", router) )
}
