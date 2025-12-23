package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("index.html")
	htmlBody := strings.TrimRight(string(content), "\n")

	baseURL, _ := url.Parse("http://blog.boot.dev")
	urls, _ := getURLsFromHTML(htmlBody, baseURL)

	fmt.Printf("URLs: %v\n", urls)
}
