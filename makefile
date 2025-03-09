.PHONY: run build docker-up docker-down test

VERSION=1.0.0

IMAGE_NAME=myapp

run:
	APP_ENV=dev go run src/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o app src/main.go

docker-build:
	docker build -t $(IMAGE_NAME):$(VERSION) .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# test:
# go test -v ./...

