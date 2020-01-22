

// *** Steps to completing code challenge ***

// 1.  Set up a simple Go server on port 8081
//     - Write a GET request that prints to page to test the server is working

// 2. Create a POST route that takes in a "count"
//     - Write a simple test function that gets passed into the POST route.
//         - For initial testing purposes, use the function to simply print the "count"

// 3. Add a 1 second sleep timer to the function so that it pauses for a second before it prints the "count"

// 4. A success message should be given once the function has completed running.

// At this point the basic structure of the server and challenge should be completed. Once this
// is running and has been tested we can move into figuring out the details of the challenge.


package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type myData struct {
	Count int
}

func timer(iChan chan int, val int, i int) {
	defer wg.Done()
	iString := strconv.Itoa(i)
	fmt.Println("Timer Function " + iString + " Started")
	
	time.Sleep(time.Second)
	fmt.Println("A second has passed for function # " + iString + ". We can now pass " + iString + " into the channel to be summed")
	iChan <- i
}

func sumChanValues(a int,b int) {
	sum := a + b
	aString := strconv.Itoa(a)
	bString := strconv.Itoa(b)
	sumString := strconv.Itoa(sum)
	fmt.Println("The sum of " + aString + " + " + bString + " = " + sumString)

}

// This route was used to get the server up and running and test that it is working correctly.
func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage is working!")
}

func testPost(w http.ResponseWriter, req *http.Request) {
	
	// Grab the POST json data
	decoder := json.NewDecoder(req.Body)

	var data myData

	// Check that the passed in data matches the type structure of the data I am expecting
	// If data is not correct, throw an error
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	// Get the count value that is being passed in
	count := data.Count

	// Create the channel for passing data through
	// We pass count into the channel to handle buffering
	iVal := make(chan int, count)
	
	// We loop over the count. wg.Add() is ticking through the buffering
	// The timer functions main goal is to pass each looping value of i into the iVal channel to be used later
	fmt.Println("")
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go timer(iVal, count, i)
	}

	// wg.Wait() is simply saying don't close the channel until all the values have been passed into the channel
	wg.Wait()
	close(iVal)

	// Once the channel is closed we can iterate through it and grab each of the values that were stored in it. 
	fmt.Println("")
	for item := range iVal {
		go sumChanValues(item, count)
	}

	// Return text back from the request saying the request was completed successfully
	fmt.Fprint(w, "You have succesfully completed the POST request!")
	
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	// Route used for testing that server is working
	myRouter.HandleFunc("/", homePage).Methods("POST")


	myRouter.HandleFunc("/post", testPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequests()
}