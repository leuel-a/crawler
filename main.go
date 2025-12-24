package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	}

	BASE_URL := args[0]
	html, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Printf("unable to fetch the page content for url: %s\n", BASE_URL)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", BASE_URL)
	fmt.Printf("html: %s\n", html)
}
