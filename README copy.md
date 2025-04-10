# PromQL Exporter
> A Prometheus Exporter for Coralogix Metrics API

This exporter fetches metrics from Coralogix's PromQL-compatible API endpoint and exposes them in Prometheus format.

## 🚀 Run

Using Docker:

```sh
docker run --rm -it \
  -p 9517:9517 \
  -e ENDPOINT=https://ng-api-http.coralogix.com \
  -e CX_API_KEY=<your-api-key> \
  -e METRICS="up,container_cpu_time_s_total" \
  ryantanjunming/promql_exporter
```

Or using a binary:

```sh
./promql_exporter \
  --endpoint https://ng-api-http.coralogixsg.com \
  --cx-api-key your-api-key \
  --metrics "up,container_cpu_time_s_total"
```

## 💡 Usage

```sh
./promql_exporter
usage: promql_exporter --endpoint=https://ng-api-http.coralogixsg.com [<flags>]

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
- `--endpoint`: Coralogix API endpoint (e.g., https://ng-api-http.coralogixsg.com)
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

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```
