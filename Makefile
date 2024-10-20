.PHONY: test coverage clean

test:
	go test ./...
	go test ./... -coverprofile=coverage.out

coverage: test
	go tool cover -html=coverage.out
	rm -f coverage.out

clean:
	rm -f coverage.out