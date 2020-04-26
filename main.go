package main

import (
	"fmt"
	"net/http"
	"os"
)

var count = 0

func NewHome(name string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		count += 1
		if count >= 5 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "You got it on %s, %d.\n", name, count)
			return
		}

		fmt.Printf("Receive Request from %s on %s\n", r.Host, name)
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "You've hit %s on %s\n", hostname, name)
		return
	}
}

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", NewHome("mux1"))
	ch := make(chan struct{})

	http.HandleFunc("/", NewHome("default mux"))
	go http.ListenAndServe("0.0.0.0:8080", nil)
	go http.ListenAndServe("0.0.0.0:8090", mux1)

	<-ch
}
