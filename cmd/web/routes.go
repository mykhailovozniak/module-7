package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	helloHandler := http.HandlerFunc(app.helloHandler)
	materialsHandler := http.HandlerFunc(app.materialsHandler)
	postHandler := http.HandlerFunc(app.postHandler)

	mux.Handle("/hello", correlationId(helloHandler))
	mux.Handle("/materials", correlationId(materialsHandler))
	mux.Handle("/post", correlationId(postHandler))

	return mux
}
