package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	testCases := []struct {
		name    string
		pageURL string
		html    string
		want    PageData
	}{
		{
			name:    "basic: h1, main paragraph, relative link and img",
			pageURL: "https://blog.boot.dev",
			html: `
<html>
  <body>
    <h1>Hello World</h1>
    <main><p>First paragraph inside main.</p></main>
    <a href="/about">About</a>
    <img src="/logo.png" alt="Logo">
  </body>
</html>`,
			want: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Hello World",
				FirstParagraph: "First paragraph inside main.",
				OutgoingLinks:  []string{"https://blog.boot.dev/about"},
				ImageURLs:      []string{"https://blog.boot.dev/logo.png"},
			},
		},
		{
			name:    "fallback paragraph when no <main>",
			pageURL: "https://blog.boot.dev",
			html: `
<html>
  <body>
    <h1>Title</h1>
    <p>Outside paragraph wins.</p>
    <a href="/x">x</a>
    <img src="/img.png">
  </body>
</html>`,
			want: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Title",
				FirstParagraph: "Outside paragraph wins.",
				OutgoingLinks:  []string{"https://blog.boot.dev/x"},
				ImageURLs:      []string{"https://blog.boot.dev/img.png"},
			},
		},
		{
			name:    "malformed HTML still parsed; absolute link and image",
			pageURL: "https://blog.boot.dev",
			html: `
<html body>
  <h1>Messy</h1>
  <a href="https://other.com/path">Other</a>
  <img src="https://cdn.boot.dev/banner.jpg">
</html body>`,
			want: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Messy",
				FirstParagraph: "", // no <p> present
				OutgoingLinks:  []string{"https://other.com/path"},
				ImageURLs:      []string{"https://cdn.boot.dev/banner.jpg"},
			},
		},
		{
			name:    "no h1 and no paragraph",
			pageURL: "https://blog.boot.dev",
			html: `
<html>
  <body>
    <a href="/only-link">Only link</a>
    <img src="/only.png">
  </body>
</html>`,
			want: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://blog.boot.dev/only-link"},
				ImageURLs:      []string{"https://blog.boot.dev/only.png"},
			},
		},
		{
			name:    "multiple links and images preserve order",
			pageURL: "https://blog.boot.dev",
			html: `
<html><body>
  <h1>t</h1>
  <main><p>p</p></main>
  <a href="/a1">a1</a>
  <a href="https://x.dev/a2">a2</a>
  <img src="/i1.png">
  <img src="https://x.dev/i2.png">
</body></html>`,
			want: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "t",
				FirstParagraph: "p",
				OutgoingLinks: []string{
					"https://blog.boot.dev/a1",
					"https://x.dev/a2",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/i1.png",
					"https://x.dev/i2.png",
				},
			},
		},
		{
			name:    "invalid base URL → empty link/image slices",
			pageURL: `:\\invalidBaseURL`,
			html: `
<html>
  <body>
    <h1>Title</h1>
    <p>Paragraph</p>
    <a href="/path">path</a>
    <img src="/logo.png">
  </body>
</html>`,
			want: PageData{
				URL:            `:\\invalidBaseURL`,
				H1:             "Title",
				FirstParagraph: "Paragraph",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := extractPageData(testCase.html, testCase.pageURL)

			if got.URL != testCase.want.URL {
				t.Errorf("URL: want %q, got %q", testCase.want.URL, got.URL)
			}
			if got.H1 != testCase.want.H1 {
				t.Errorf("H1: want %q, got %q", testCase.want.H1, got.H1)
			}
			if got.FirstParagraph != testCase.want.FirstParagraph {
				t.Errorf("FirstParagraph: want %q, got %q", testCase.want.FirstParagraph, got.FirstParagraph)
			}
			if !reflect.DeepEqual(got.OutgoingLinks, testCase.want.OutgoingLinks) {
				t.Errorf("OutgoingLinks: want %v, got %v", testCase.want.OutgoingLinks, got.OutgoingLinks)
			}
			if !reflect.DeepEqual(got.ImageURLs, testCase.want.ImageURLs) {
				t.Errorf("ImageURLs: want %v, got %v", testCase.want.ImageURLs, got.ImageURLs)
			}
		})
	}
}
