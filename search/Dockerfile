FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY ./src ./src

RUN go mod download

RUN go build -o server ./src/main.go

CMD ["./server"]
