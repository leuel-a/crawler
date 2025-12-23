package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var images []string
	document.Find("img").Each(func(_ int, selection *goquery.Selection) {
		if src, ok := selection.Attr("src"); ok {
			parsedSrc, err := url.Parse(src)
			if err != nil {
				fmt.Printf("[Error] unable to parse the src: %v\n", parsedSrc)
				return
			}
			images = append(images, baseURL.ResolveReference(parsedSrc).String())
		}
	})

	return images, nil
}
