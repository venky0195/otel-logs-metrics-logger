# otel-logs-metrics-logger

This collector is built in Go and utilizes the OpenTelemetry framework for telemetry purposes.

## Setup

### Prerequisites

- Go installed version 1.20 >=

### Setup

Before running the service, set up the configuration file for the OpenTelemetry collector:

```bash
make setup
```

### Running the Service

To run the service, execute the following command:

```bash
make run
```

This will start the service using the configuration specified in 'otelcol-config.yaml'.
The HTTP server will be available at the endpoint `localhost:4318`
The GRPC server will be available at the endpoint `localhost:4317`
The healthcheck of this service will be available at the endpoint `localhost:13133`

### Configuration Files

- `otelcol-config.yaml`: This file is used to configure the Open Telemetry Collector itself.