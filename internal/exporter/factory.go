package exporter

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	typeStr = "exporter"
)

func NewFactory(logExporter *LogExporter, metricExporter *MetricsExporter) exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (exporter.Logs, error) {
			return createLogsExporter(ctx, set, cfg, logExporter)
		}, component.StabilityLevelAlpha),
		exporter.WithMetrics(func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (exporter.Metrics, error) {
			return createMetricsExporter(ctx, set, cfg, &MetricsExporter{})
		}, component.StabilityLevelAlpha),
		exporter.WithTraces(func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (exporter.Traces, error) {
			return createTraceExporter(ctx, set, cfg, &TracesExporter{})
		}, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
	logExporter *LogExporter,
) (exporter.Logs, error) {

	return exporterhelper.NewLogsExporter(ctx, set, cfg,
		func(ctx context.Context, ld plog.Logs) error {
			return logExporter.ConsumeLogs(ctx, ld)
		},
	)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
	metricExporter *MetricsExporter,
) (exporter.Metrics, error) {

	return exporterhelper.NewMetricsExporter(ctx, set, cfg,
		func(ctx context.Context, ld pmetric.Metrics) error {
			return metricExporter.ConsumeMetrics(ctx, ld)
		},
	)
}

func createTraceExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
	traceExporter *TracesExporter,
) (exporter.Traces, error) {

	return exporterhelper.NewTracesExporter(ctx, set, cfg,
		func(ctx context.Context, td ptrace.Traces) error {
			return traceExporter.ConsumeTraces(ctx, td)
		},
	)
}
