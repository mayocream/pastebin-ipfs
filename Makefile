build:
	go build -o bin/pstbin cmd/main.go

compose:
	sudo docker-compose up -d

run: compose
	go run cmd/main.go

test:
	go test -v --v ./...