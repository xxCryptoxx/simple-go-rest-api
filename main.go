package main

import (
	"fmt"
	"log"
	"net/http"
)

// this function is responsible for localhost:8000/
func testHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Home: Endpoint Hit")
	fmt.Fprintln(w, "Test Home: Endpoint Hit")
}

// this function is responsible for localhost:8000/test
func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Page: Endpoint Hit")     // Prints to terminal
	fmt.Fprintln(w, "Test Page: Endpoint Hit") // Prints to site
}

// this function is responsible handling requests
// http.HandleFunc("/DirectToEndpoint", EndpointBackend)
func requestHandler() {
	http.HandleFunc("/", testHome)
	http.HandleFunc("/test", testPage)

	// Listens and serves the port, default url is localhost
	// if you want to change it to an IP or domain name then
	// log.Fatal(http.ListenAndServe("192.168.1.1:8000", yourFunction))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// This function is called
func main() {
	fmt.Println("Running") // Prints to terminal
	requestHandler()       // Calls the request handler
}
