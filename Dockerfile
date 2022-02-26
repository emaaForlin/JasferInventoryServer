FROM golang:1.17.7 AS builder
ARG VERSION=dev
WORKDIR /go/src/app
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM debian:stretch-slim
RUN mkdir /app
COPY --from=builder /go/src/app/main /go/bin/main
ENV DB_HOST="" DB_PORT=3306 DB_USER="" DB_PASS="" DB_NAME=""
ENV PATH="/go/bin:${PATH}"
EXPOSE 9090
CMD ["main"]