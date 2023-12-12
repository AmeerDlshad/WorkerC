package main

import (
	"fmt"
	"net/http"

	workers "github.com/AmeerDlshad/WorkerC"
	"github.com/AmeerDlshad/WorkerC/cloudflare"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "MY_ENV: %s", cloudflare.Getenv(req.Context(), "MY_ENV"))
	})
	workers.Serve(handler)
}
