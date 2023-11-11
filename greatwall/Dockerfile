FROM golang:1.18-alpine

WORKDIR /app

COPY greatwall.go .

RUN go build -o greatwall-server greatwall.go

EXPOSE 8080

CMD ["/app/greatwall-server"]
