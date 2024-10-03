FROM golang:alpine AS build

WORKDIR /auth

COPY . .

RUN go mod download

RUN go build cmd/auth_server.go

FROM alpine:latest

WORKDIR /auth

EXPOSE 3080

COPY .env .

COPY --from=build /auth/auth_server /auth/auth_server

CMD ["./auth_server"]