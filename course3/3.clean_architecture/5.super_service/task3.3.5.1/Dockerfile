FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY .env .

EXPOSE 8080

COPY --from=build /app/main /app/main

CMD ["./main"]