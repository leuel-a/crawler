package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Printf("[Error] unable to load document with goquery `%v`\n", err)
		return nil, nil
	}

	urls := []string{}
	document.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			parsedHref, err := url.Parse(href)
			if err != nil {
				fmt.Printf("[Error] unable to parse the href: %v\n", href)
				return
			}
			urls = append(urls, baseURL.ResolveReference(parsedHref).String())
		}
	})

	return urls, nil
}
