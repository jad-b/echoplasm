package main

import (
	"flag"
	"net/http"

	"github.com/jad-b/echoplasm/client"
)


// EchoHandler writes the request body into the response.
func EchoHandler(w http.ResponseWriter, req *http.Request) {
	err := req.Write(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	// Flag parsing
	daemonPtr := flag.Bool("daemon", false,
	"Whether or not to leave the server in the background")
    flag.Parse()

	http.HandleFunc("/", EchoHandler)
	if *daemonPtr {
		http.ListenAndServe(":8080", nil)
	} else {
		go http.ListenAndServe(":8080", nil)
		client.SamplePOSTCalls()
	}
}
