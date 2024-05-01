default: test

.PHONY: test
test:
	go clean -testcache
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out