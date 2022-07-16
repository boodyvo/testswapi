FROM golang:1.18-alpine as builder

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o app cmd/swapi/main.go

FROM alpine:3

COPY --from=builder /app/app /usr/local/bin/app

CMD ["app"]
