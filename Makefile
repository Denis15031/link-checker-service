.PHONY: build run docker-build docker-run test clean

build:
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

docker-build:
	docker build -t link-checker-service .

docker-run:
	docker run -p 8080:8080 link-checker-service

test:
	go test ./...

clean:
	rm -f bin/server