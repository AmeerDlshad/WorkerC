package main

import (
	"net/http"

	workers "github.com/AmeerDlshad/WorkerC"
	"github.com/AmeerDlshad/WorkerC/_examples/d1-blog-server/app"
)

func main() {
	http.Handle("/articles", app.NewArticleHandler())
	workers.Serve(nil) // use http.DefaultServeMux
}
