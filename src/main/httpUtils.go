package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getJsonFromUrl(url string) string {
	res, e := http.Get("https://svc.metrotransit.org/NexTrip/Routes?format=json")

	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(1)
	}

	resBody, e := ioutil.ReadAll(res.Body)

	if e != nil {
		log.Fatal(e)
	}

	return (string(resBody))
}
