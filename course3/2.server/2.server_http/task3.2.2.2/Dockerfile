FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY .env .

COPY --from=build /app/main /app/main

CMD ["./main"]