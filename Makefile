GO = go
OUTPUT = crawler
WEB_URL = https://wagslane.dev

build:
	@$(GO) build -o $(OUTPUT) && ./$(OUTPUT) $(WEB_URL)

test:
	$(GO) test ./...
