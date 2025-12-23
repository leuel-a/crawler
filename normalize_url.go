package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeUrl(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)

	if err != nil {
		fmt.Printf("[Error] normalizing the rawURL %s, error: %v\n", rawURL, err)
		return "", err
	}

	parsedURL := fmt.Sprintf("%s%s", u.Host, strings.TrimRight(u.Path, "/"))
	return parsedURL, nil
}
