FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gov-osint-tool ./cmd/gov-osint-tool

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/gov-osint-tool .
COPY webui ./webui
EXPOSE 8080
CMD ["./gov-osint-tool"]
