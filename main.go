package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Panic("Error when starting the http server", err)
	}
}
