package exporter

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricsExporter is a simple exporter that logs received metrics.
type MetricsExporter struct{}

// NewMetricsExporter creates a new MetricsExporter.
func NewMetricsExporter() *MetricsExporter {
	return &MetricsExporter{}
}

// ConsumeMetrics logs the received metrics.
func (me *MetricsExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	// Use the OpenTelemetry-provided JSON marshaler
	marshaler := pmetric.JSONMarshaler{}

	// Marshal the metrics to JSON
	jsonData, err := marshaler.MarshalMetrics(md)
	if err != nil {
		log.Printf("Error marshaling metrics: %v", err)
		return err
	}

	// Log the JSON representation of the metrics
	log.Printf("Metrics JSON: %s", string(jsonData))
	return nil
}
