# echoplasm
Echo server for examining HTTP requests/responses.

## Usage
1. Create some new client code under `client/`. For instance, example POST calls go in `client/post.go`.
1. Hook them into the server with something like:
```Go
// In `github.com/jad-b/echoplasm/main.go`:

// ...

func main()
    // ...
    go http.ListenAndServe(":8080", nil)
    client.SamplePOSTCalls()
    client.ExampleGETCall1()
    client.DemonstratoryPUTting()
}
```
