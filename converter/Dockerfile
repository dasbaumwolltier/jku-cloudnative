FROM golang:1.19-alpine AS builder

COPY . /app
WORKDIR /app

RUN go build -o /app/converter

# Path: converter/Dockerfile
FROM alpine:3.14

EXPOSE 8080

COPY --from=builder /app/converter /converter
RUN chmod +x /converter

ENTRYPOINT ["/converter"]