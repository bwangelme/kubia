package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var count = 0

func Home(w http.ResponseWriter, r *http.Request) {
	count += 1
	if count >= 5 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "You got it, %d.\n", count)
		return
	}

	fmt.Printf("Receive Request from %s\n", r.Host)
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "You've hit %s\n", hostname)
	return
}

func main() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
