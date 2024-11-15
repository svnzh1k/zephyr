FROM golang:1.23 AS builder

WORKDIR /gp/src/app

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]