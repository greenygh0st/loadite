package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	uriPtr := flag.String("u", "", "URI to whack (Required)")
	cPtr := flag.Int("c", 1000, "Number of requests to make")
	jwtPtr := flag.String("jwt", "", "Added a bearer token to the request")

	flag.Parse()

	if *uriPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// set wait group count
	wg.Add(*cPtr)

	fmt.Println("🏓 Stand by to beat on: ", *uriPtr)
	// fmt.Println("done.")

	for i := 0; i < *cPtr; i++ {
		// Actually make the request
		go makeRequest(*uriPtr, *jwtPtr)
		fmt.Println("Started request " + strconv.Itoa(i+1) + " of " + strconv.Itoa(*cPtr))
	}

	fmt.Println("Finished the pounding 🎉")
}

func makeRequest(uri string, bearer string) {
	// Create a new request using http
	// TODO: need to be able to send a method
	req, _ := http.NewRequest("GET", uri, nil)

	if bearer != "" {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + bearer

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	// Send req using http Client
	client := &http.Client{}
	// resp, _ := client.Do(req) //
	if resp, err := client.Do(req); err != nil {
		log.Println(err)
	} else {
		// return readBody(resp.Body)
		log.Println(resp.Status)
		log.Println(resp)
		defer resp.Body.Close()
	}
	// we finished this particular whack
	wg.Done()
}
