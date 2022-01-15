FROM golang:1.16-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main ./app/


# stage 2
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder ./main .
EXPOSE 2801
CMD ["./main"]