package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove end forward slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove double end forward slash",
			inputURL: "https://blog.boot.dev/path//",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove end forward slash in subpath",
			inputURL: "https://blog.boot.dev/path/subpath/",
			expected: "blog.boot.dev/path/subpath",
		},
	}

	for index, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := normalizeUrl(testCase.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", index, testCase.name, err)
				return
			}
			if actual != testCase.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", index, testCase.name, testCase.expected, actual)
			}
		})
	}
}
