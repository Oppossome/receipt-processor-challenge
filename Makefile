.PHONY: run-tools
run-tools:
	go generate ./tools/tools.go

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run
run: tidy
	go run ./cmd/receipt-processor

.PHONY: test
test: tidy
	go test ./...