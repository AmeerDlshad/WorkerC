package main

import (
	"fmt"
	"io"
	"net/http"

	workers "github.com/AmeerDlshad/WorkerC"
	"github.com/AmeerDlshad/WorkerC/cloudflare"
	"github.com/AmeerDlshad/WorkerC/cloudflare/fetch"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		bind := cloudflare.GetBinding(ctx, "hello")
		fc := fetch.NewClient(fetch.WithBinding(bind))

		hc := fc.HTTPClient()
		res, err := hc.Do(req)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		io.Copy(w, res.Body)
	})
	workers.Serve(handler)
}
