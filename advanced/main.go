
// ***** steps.txt contains my thought process while working through this challenge *****


package main

import (
	"fmt"
	"log"
	"net/http"
)

// This route was used to get the server up and running and test that it is working correctly.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage is working!")
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}