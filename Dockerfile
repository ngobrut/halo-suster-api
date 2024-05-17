FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]