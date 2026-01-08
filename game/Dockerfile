FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o game ./cmd

CMD ["./game"]