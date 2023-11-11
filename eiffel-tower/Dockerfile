FROM golang:1.18-alpine

WORKDIR /app

COPY eiffel.go .

RUN go build -o eiffel-server eiffel.go

EXPOSE 8080

CMD ["/app/eiffel-server"]
