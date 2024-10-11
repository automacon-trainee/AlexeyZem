FROM golang:alpine AS build

WORKDIR /user

COPY . .

RUN go mod download

RUN go build cmd/user_server.go

FROM alpine:latest

WORKDIR /user

COPY .env .

EXPOSE 1080

COPY --from=build /user/user_server /user/user_server

CMD ["./user_server"]