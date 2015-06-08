package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURL   = "http://127.0.0.1:8080"
	echo      = "/echo/"
	bodyBreak = `
###############################################################################

`
)

func printResponseBody(respBody io.ReadCloser) {
	defer respBody.Close()
	body, _ := ioutil.ReadAll(respBody)
	log.Printf("Request body:\n%s%s", body, bodyBreak)
}

// SamplePOSTCalls demonstrates what POSTing data to a server looks like by
// writing the Reponse bodies to stdout. It expects to call against an "/echo/"
// endpoint that echoes the request into the response.
func SamplePOSTCalls() {
	resp := examplePOST()
	printResponseBody(resp.Body)

	resp = exampleFormValuePOST()
	printResponseBody(resp.Body)
}

func examplePOST() (resp *http.Response) {
	log.Print("POSTing example data")
	buf := bytes.NewBufferString("Hey! I'm example data in a POST")
	resp, _ = http.Post(baseURL+echo, "text/text", buf)
	return
}

func exampleFormValuePOST() (resp *http.Response) {
	log.Print("POSTing example form data")
	vals := url.Values{"Key1": {"Value1"}, "Key2": {"Value2", "Value3"}}
	resp, _ = http.PostForm(baseURL+echo, vals)
	return
}
