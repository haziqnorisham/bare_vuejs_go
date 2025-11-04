package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("dist"))
	http.Handle("/", fileServer)

	log.Printf("Listening to localhost:%s...", "8080")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil); err != nil {
		log.Fatal(err.Error())
	}
}
