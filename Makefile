run:
	go run main.go

test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: run test