build:
	go build -o http_server ./cmd/app/main.go

run: build
	./http_server

docker-build:
	GO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./out/app ./cmd/app