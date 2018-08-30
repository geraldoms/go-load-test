package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	reqs := 300				// Total of requests for the load test
	start := time.Now()		// Get the current time for 
	ch := make(chan string, reqs)	// Channel to get info from the goroutines
	counter := 0			// error counter

	// This loop creates a new goroutine for each request to the server.
	// The func uses a counter (by reference) to get the errors and a channel to communicate
	for i := 1; i <= reqs; i++ {
		go DoGet("http://localhost:3000", ch, i, &counter)
		//time.Sleep(time.Millisecond) 		// Sleeping time to send requests between time intervals
	}

	// Read all messages from the channel and print on the screen
	for i := 1; i <= reqs; i++ {
		fmt.Println(<-ch)
	}

	// Shows the values of the tests
	fmt.Printf("\n\nTotal of requests: [%d], time: [%.4f secs].\n", reqs, time.Since(start).Seconds())
	fmt.Printf("Success: %d - Errors: %d\n", reqs - counter, counter)
}

// This func sends a get request to the server and write the results into the channel
func DoGet(url string, ch chan<- string, idx int, counter *int) {
	start := time.Now()
	resp, err := http.Get(url)
	status := "ERROR"
	if err == nil {
		status = resp.Status
		resp.Body.Close()
	} else {
		*counter++
	}
	ch <- fmt.Sprintf("Idx: [%d], Status: [%s], Time: [%f secs].", idx, status, time.Since(start).Seconds())
}
