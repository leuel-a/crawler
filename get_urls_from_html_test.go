package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLBasic(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="https://blog.boot.dev"></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldb't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
				<a href="/home""></a>
				<a href="/about""></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev/home", "https://blog.boot.dev/about"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLPathRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
				<a href="posts/home""></a>
				<a href="posts/about""></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev/posts/home", "https://blog.boot.dev/posts/about"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}

}

func TestGetURLsFromHTMLProtocolRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
				<a href="//blog.boot.dev/contact"></a>
				<a href="//blog.boot.dev/posts/about""></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev/contact", "https://blog.boot.dev/posts/about"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}

}

func TestGetURLsFromHTMLFragments(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
				<a href="#contact"></a>
				<a href="#section""></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev#contact", "https://blog.boot.dev#section"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLSearchQuery(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
				<a href="/search?q=label"></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{"https://blog.boot.dev/search?q=label"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLSkipInvalidHrefs(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
                <a href=":invalid:url"></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLMissingAttributes(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody :=
		`
		<html>
			<body>
                <a>no href</a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v. got %v", expected, actual)
	}
}
