build: 
		@go build -o bin/golang

run: build
		@./bin/golang

test: 
		@go test -v ./...