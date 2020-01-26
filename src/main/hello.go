package main

import (
	"fmt"
)

func main() {
	fmt.Printf(hello() + "\n\n")

	getRoutes()
	fmt.Printf("\n\n")
	getProviders()
}

func hello() string {
	return "Hello, world\n"
}
