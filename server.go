package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	go func () {
			log.Fatal(http.ListenAndServe(":8080", r))
	}()

	r.HandleFunc("/", handleRequest)
	r.HandleFunc("/admin", handleRequest)
	r.HandleFunc("/root", handleRequest)
	r.HandleFunc("/test", handleRequest)

	http.ListenAndServe(":80", r)
}


func handleRequest(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		method := req.Method
		fmt.Println(path)

		fmt.Fprintf(w, "You've made a %s request for path %s\n", method, path)
	}
