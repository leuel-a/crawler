package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("[Error] unable to load document with goquery `%v`\n", err)
		return ""
	}

	var headingText = doc.Find("h1").First().Text()
	return strings.TrimSpace(headingText)
}

func getFirstParagraphFromHTML(html string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("[Error] unable to load document with goquery `%v`\n", err)
		return ""
	}

	main := document.Find("main")
	paragraphText := ""

	if main.Length() > 0 {
		paragraphText = main.Find("p").First().Text()
	} else {
		paragraphText = document.Find("p").First().Text()
	}

	return strings.TrimSpace(paragraphText)
}
