.PHONY: build doc fmt lint run test vet


default: build

# Build all packages
build:
	go build

run:
	go build && ./webapp

test:
	go test -v ./...

# Start documentation server with search
doc:
	godoc -http=:8080 -index

vet:
	go vet
