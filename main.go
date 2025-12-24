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
	pages := map[string]int{}

	crawlPage(BASE_URL, BASE_URL, pages)
}
