build:
	go build -o bin/pesakit cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

sync:
	go mod tidy
	go mod download
	go fmt
	git add -A