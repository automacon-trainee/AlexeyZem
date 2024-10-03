FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/user_server.go

FROM alpine:latest

WORKDIR /app

COPY .env .

EXPOSE 8080

COPY --from=build /app/user_server /app/user_server

CMD ["./user_server"]