FROM golang:alpine AS build

WORKDIR /geo

COPY . .

RUN go mod download

RUN go build cmd/geodata_server.go

FROM alpine:latest

WORKDIR /geo

EXPOSE 2080

COPY .env .

COPY --from=build /geo/geodata_server /geo/geodata_server

CMD ["./geodata_server"]