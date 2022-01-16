FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

# stage 2
FROM alpine:3.14
WORKDIR /root
COPY --from=builder /app/main . 
EXPOSE 2801
CMD ["/root/main"]