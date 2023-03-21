FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY ./src ./src

RUN go mod download

RUN go build -o main ./src/main.go

CMD ["./main"]
