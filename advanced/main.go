
// ***** steps.txt contains my thought process while working through this challenge *****


package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
)

type myData struct {
	Count int
}

func timer(val int) {
	fmt.Println("Timer Function Started")
	for i := 1; i <= val; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	fmt.Println("SUCCESS!")
}

// This route was used to get the server up and running and test that it is working correctly.
func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage is working!")
}

func testPost(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	count := data.Count

	timer(count)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/post", testPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequests()
}