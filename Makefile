build:
	go build -o ./bin/projectx

run: build
	./bin/projectx

test:
	go mod tidy
	go test ./...
	