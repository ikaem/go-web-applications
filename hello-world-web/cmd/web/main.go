package main

import (
	"fmt"
	"net/http"

	"github.com/ikaem/hello-world-web/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting app on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil) // -> this does not work

}
