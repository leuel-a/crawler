GO = go
GO_MAIN = main

build:
	@$(GO) build -o $(GO_MAIN).out .
	@./$(GO_MAIN).out

test:
	$(GO) test ./...
