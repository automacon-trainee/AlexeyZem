swagger:
	swag init -g cmd/app/main.go

test:
	TZ=UTC go test ./... -v

lint:
	golangci-lint run -c .golangci.yaml
