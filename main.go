package main

import (
	"fmt"
	"net/http"
)

// https://learn.microsoft.com/en-us/azure/app-service/quickstart-golang

func main() {
	http.HandleFunc("/", HelloServer)
	//println("serving http://localhost:8080/world")
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
