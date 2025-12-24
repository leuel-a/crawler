# Go Web Crawler

A simple web crawler written in Go that fetches the HTML content of a given URL.

## Usage

To run the crawler, provide a starting URL as a command-line argument.

```sh
go run main.go <your-starting-url>
```

**Example:**
```sh
go run main.go https://example.com
```

## Development

### Building

You can build the project using the provided Makefile:

```sh
make build
```
This will create an executable named `crawler` and run it against a default URL.

### Testing

To run the test suite:

```sh
make test
```
