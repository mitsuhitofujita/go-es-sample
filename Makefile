
.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run sample.go
