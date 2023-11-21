package main

import "net/http"

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	http.HandleFunc("/snippet", app.showSnippet)
	http.HandleFunc("/snippet/create", app.createSnippet)

	return mux
}
