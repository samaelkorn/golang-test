package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	result := user.getUser()
	fmt.Fprintf(w, result)
}
