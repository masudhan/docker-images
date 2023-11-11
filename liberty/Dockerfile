FROM golang:1.18-alpine

WORKDIR /app

COPY liberty.go .

RUN go build -o liberty-server liberty.go

EXPOSE 8080

CMD ["/app/liberty-server"]
