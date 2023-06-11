.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: bench
bench: ## Run benchmark test. See https://pkg.go.dev/cmd/go#hdr-Testing_flags
	go test ./... -bench . -benchtime 5s -timeout 0 -run=XXX -cpu 1 -benchmem