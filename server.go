package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Define the path and the handler function
	http.HandleFunc("/", handler)
	// listen and serve on port 3000
	http.ListenAndServe(":3000", nil)
}

// The func just pipe a Hello! into the response 
func handler(w http.ResponseWriter, r *http.Request) {
	// Sleeping time representing a small workload
	time.Sleep(100 * time.Millisecond)
	fmt.Fprint(w, "<h1>Hello!</h1>")
}