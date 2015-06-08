package main

import (
	"net/http"

	"./client"
)

// EchoHandler writes the request body into the response.
func EchoHandler(w http.ResponseWriter, req *http.Request) {
	err := req.Write(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", EchoHandler)
	go http.ListenAndServe(":8080", nil)
	client.SamplePOSTCalls()
}
