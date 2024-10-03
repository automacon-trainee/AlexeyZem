FROM golang:alpine AS build

WORKDIR /notify

COPY . .

RUN go mod download

RUN go build cmd/notify_server.go

FROM alpine:latest

WORKDIR /notify

EXPOSE 4080

COPY .env .

COPY --from=build /notify/notify_server /notify/notify_server

CMD ["./notify_server"]