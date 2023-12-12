package main

import (
	"io"
	"net/http"

	workers "github.com/AmeerDlshad/WorkerC"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello!"
		w.Write([]byte(msg))
	})
	http.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, req.Body)
	})
	workers.Serve(nil) // use http.DefaultServeMux
}
