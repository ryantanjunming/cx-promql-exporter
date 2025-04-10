
FROM alpine:latest as certs

RUN apk --update add ca-certificates

COPY . .

ENTRYPOINT ["/cx-promql-exporter"]

EXPOSE 9517
