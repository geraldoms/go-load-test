# Go Load Test

Basic HTTP load test using Golang. In its core, Go provides an HTTP server and, by default, each request starts a new goroutine (lightweight thread of execution), then the requests are processed concurrently allowing high-performance in web applications.

## Requirements
* Go 1.10 or later

## Usage 

`$ go run server.go`

`$ go run client.go`
