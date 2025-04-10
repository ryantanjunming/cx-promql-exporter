# PromQL Exporter
> A Prometheus Exporter for PromQL

This exporter fetches metrics from PromQL-compatible API endpoint and exposes them in Prometheus format.

Added functionality for querying Coralogix's Metric API which is used as a datasource for Grafana Plugin.

More Info: 
- <https://coralogix.com/docs/user-guides/data-query/metrics-api/#promql-compatible-query-data-source>
- <https://coralogix.com/docs/user-guides/visualizations/grafana-plugin/>

Inspired by: <https://github.com/samber/promql-exporter>

## 🏗️ Build

To build the Docker image locally:

```sh
# From the project root directory
docker build -t promql-exporter .
```

The Dockerfile uses a multi-stage build process:
1. First stage builds the Go binary using `golang:1.21-alpine`
2. Final stage creates a minimal image using `alpine:3.19`

The resulting image is optimized for size and security, with:
- Minimal dependencies
- CGO disabled for better compatibility
- Non-root user execution
- Port 9517 exposed for metrics

To build the binary directly:

```sh
# Install dependencies
go mod download

# Build the binary
go build -o promql_exporter

# For production builds with optimizations
CGO_ENABLED=0 GOOS=linux go build -o promql_exporter
```

## 🚀 Run

Using Docker:

```sh
docker run --rm -it \
  -p 9517:9517 \
  -e ENDPOINT=https://ng-api-http.coralogix.com \
  -e CX_API_KEY=<cx-api-key> \
  -e METRICS="<metrics-comma-separated>" \
  your-container-repo/cx-promql-exporter
```

Or using a binary:

```sh
./promql_exporter \
  --endpoint https://ng-api-http.coralogix.com \
  --cx-api-key <cx-api-key> \
  --metrics "up,container_cpu_time_s_total"
```

## 💡 Usage

```sh
./promql_exporter
usage: promql_exporter --endpoint=https://ng-api-http.coralogix.com [<flags>]

Flags:
  -h, --help                           Show context-sensitive help
      --endpoint                       Coralogix API endpoint ($ENDPOINT)
      --cx-api-key                     Coralogix API Key ($CX_API_KEY)
      --header                         Additional HTTP headers ($HEADER)
      --metrics                        Comma-separated list of metrics to query ($METRICS)
      --namespace                      Namespace for metrics ($PROMQL_EXPORTER_NAMESPACE)
      --web.listen-address=":9517"     Address to listen on for web interface and telemetry
      --web.telemetry-path="/metrics"  Path under which to expose metrics
      --log.format="txt"               Log format (txt or json)
      --version                        Show application version
```

## 🔧 Configuration

### Required Parameters
- `--endpoint`: Coralogix API endpoint (e.g., https://ng-api-http.coralogix.com)
- `--cx-api-key`: Coralogix API key for authentication

### Optional Parameters
- `--metrics`: Comma-separated list of metrics to query (e.g., "up,container_cpu_time_s_total")
- `--namespace`: Prefix for all exported metrics
- `--header`: Additional HTTP headers (can be specified multiple times)
- `--web.listen-address`: Port to expose metrics on (default: :9517)
- `--web.telemetry-path`: Path to expose metrics on (default: /metrics)

## 📝 License

Copyright © 2024 [Ryan Tan](https://github.com/ryantanjunming).

This project is [MIT](./LICENSE) licensed.
```
