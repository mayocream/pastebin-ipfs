all: build build-web

build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/pstbin cmd/main.go

compose:
	sudo docker-compose up -d

run: compose
	go run cmd/main.go

test:
	go test -v --v ./...

build-web:
	pnpm run build

web-live:
	pnpm run dev