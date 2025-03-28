build:
	@go build -o bin/mentee ./cmd/mentee/main.go

run: build
	@./bin/mentee --config=./config/local.yaml

test:
	@go test -v ./...

migrate:
	@go run ./cmd/migrate/main.go --config=./config/local.yaml --migrations-path=./migrations
