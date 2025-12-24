package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		fmt.Printf("[Error] unable to create request: %s\n", err)
		return "", err
	}

	request.Header.Set("User-Agent", "BootCrawler/1.0")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("[Error] error making http request: %v\n", err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		err = fmt.Errorf("[Error] got response with status code: %v\n", response.StatusCode)
		return "", err
	}

	if !strings.HasPrefix(response.Header.Get("Content-Type"), "text/html") {
		err = fmt.Errorf("[Error] response is not an html content, instead: %s\n", response.Header.Get("Content-Type"))
		return "", err
	} 

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[Error] could not read response body: %s\n", responseBody)
		return "", err
	}

	return string(responseBody), nil
}
