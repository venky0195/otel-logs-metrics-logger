package main

import (
	"github.com/rs/zerolog/log"
	loggingexporter "github.com/venky0195/otel-logs-metrics-logger/internal/exporter"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"
)

func main() {
	info := component.BuildInfo{
		Command:     "otel-logs-metrics-logger",
		Description: "Local OpenTelemetry Collector binary",
		Version:     "1.0.0",
	}

	logExporter := loggingexporter.NewLogExporter()
	metricExporter := loggingexporter.NewMetricsExporter()

	factories, err := components(logExporter, metricExporter)
	if err != nil {
		log.Fatal().Err(err)
	}

	settings := otelcol.CollectorSettings{
		BuildInfo: info,
		Factories: func() (otelcol.Factories, error) {
			return factories, nil
		},
	}

	if runInteractiveErr := runInteractive(settings); err != nil {
		log.Fatal().Err(runInteractiveErr)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := otelcol.NewCommand(params)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("collector server run finished with error:")
	}

	return nil
}
