package exporter

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/pdata/plog"
)

// LogExporter is a simple exporter that logs received logs.
type LogExporter struct{}

// NewLogExporter creates a new LogExporter.
func NewLogExporter() *LogExporter {
	return &LogExporter{}
}

// ConsumeLogs logs the received logs.
func (le *LogExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	// Use the OpenTelemetry-provided JSON marshaler
	marshaler := plog.JSONMarshaler{}

	// Marshal the logs to JSON
	jsonData, err := marshaler.MarshalLogs(ld)
	if err != nil {
		log.Printf("Error marshaling logs: %v", err)
		return err
	}

	// Log the JSON representation of the logs
	log.Printf("Logs JSON: %s", string(jsonData))
	return nil
}
