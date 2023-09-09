FROM golang:1.19 as builder

WORKDIR /usr/app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

FROM alpine

COPY --from=builder /usr/app /app

WORKDIR /app

EXPOSE 3000

CMD ["/app/server"]