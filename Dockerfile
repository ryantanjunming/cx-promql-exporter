# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o promql_exporter

# Final stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/promql_exporter /app/

EXPOSE 9517

ENTRYPOINT ["/app/promql_exporter"]
