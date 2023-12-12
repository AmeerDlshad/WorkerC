package main

import (
	"fmt"
	"net/http"

	workers "github.com/AmeerDlshad/WorkerC"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})
	workers.Serve(handler)
}
