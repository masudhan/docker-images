FROM golang:1.18-alpine

WORKDIR /app

COPY tajmahal.go .

RUN go build -o tajmahal-server tajmahal.go

EXPOSE 8080

CMD ["/app/tajmahal-server"]
