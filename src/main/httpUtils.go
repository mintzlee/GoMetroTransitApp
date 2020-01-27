package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// returns request body as a string. May return empty string.
func getDataFromUrl(url string) (string, error) {
	res, e := http.Get(url)

	if e != nil {
		fmt.Printf(e.Error())
		return "", e
	}

	resBody, e := ioutil.ReadAll(res.Body)

	if e != nil {
		log.Fatal(e)
		return "", e
	}

	return string(resBody), nil
}
