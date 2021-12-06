FROM gcr.io/distroless/static AS base

WORKDIR /
LABEL maintainer="Sergio Villanueva <sergiovillanueva@protonmail.com>"
EXPOSE 8080

FROM golang:1.17-alpine AS builder

WORKDIR /go/src/github.com/hexdump95/quadratic/
COPY . .
ENV CGO_ENABLED=0
RUN go build -o app

FROM base

COPY --from=builder /go/src/github.com/hexdump95/quadratic/app .
CMD ["/app"]
