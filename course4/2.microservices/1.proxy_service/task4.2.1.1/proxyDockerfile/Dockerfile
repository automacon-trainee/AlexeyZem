FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/proxy/proxy.go

FROM alpine:latest

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/proxy /app/proxy

CMD ["./proxy"]