FROM golang:1.26.5
WORKDIR /weatherBot
COPY . .
RUN go build -o main cmd/main.go
CMD ["./main"]