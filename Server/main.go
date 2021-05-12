package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/SIS":
		codename := req.URL.Query().Get("codename")
		fmt.Fprint(w, "Hello Agent %s\n\n\n", hello(codename))
	case "/b":
		fmt.Fprint(w, "You asked for /b")
	case "/c":
		fmt.Fprint(w, "You requested /c")
	default:
		http.Error(w, "Are you out of your mind??!?!?!", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)

}
