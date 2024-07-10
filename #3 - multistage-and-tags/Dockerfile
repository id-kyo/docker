# Stage 1: Build
FROM golang:1.18.0-bullseye AS builder
WORKDIR /app
COPY /app.go .
RUN go mod init main && CGO_ENABLED=0 go build

# Stage 2: Stage
FROM scratch AS stage
EXPOSE 80
COPY --from=builder /app/main /go/main
CMD ["/go/main"]