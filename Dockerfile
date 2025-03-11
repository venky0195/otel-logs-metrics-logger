FROM golang:1.21-alpine as build
RUN apk --no-cache add make
ARG OTEL_BINARY_NAME=otel-logs-metrics-logger
WORKDIR /app
COPY . .
RUN apk --no-cache add --update bash  # Install bash if not present
RUN make setup
RUN make build

FROM alpine:3.14
RUN apk --no-cache add curl
COPY --from=build /app/bin/${OTEL_BINARY_NAME} ./${OTEL_BINARY_NAME}
COPY --from=build /app/otelcol-config.yaml ./otelcol-config.yaml
CMD ["./otel-logs-metrics-logger", "--config=otelcol-config.yaml"]