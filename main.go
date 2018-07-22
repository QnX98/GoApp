package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	// setting headers
	w.Header().Set("Content-Type", "text/html")
	// this gets the URL path
	path := r.URL.Path
	if path == "/" {
		fmt.Fprint(w, "<h1>Fuck this shit server you at homepage niggaa.</h1>")
	} else if path == "/contact" {
		fmt.Fprint(w, "<h1>Contact bitch, Fuck this shit server you at contact page niggaa.</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 bitch, Fuck this shit server you at page not found niggaa.</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
