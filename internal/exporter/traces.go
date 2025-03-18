package exporter

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

// TracesExporter is a simple exporter that logs received traces.
type TracesExporter struct{}

// NewTracesExporter creates a new TracesExporter.
func NewTracesExporter() *TracesExporter {
	return &TracesExporter{}
}

// ConsumeTraces logs the received traces.
func (te *TracesExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// Use the OpenTelemetry-provided JSON marshaler
	marshaler := ptrace.JSONMarshaler{}

	// Marshal the traces to JSON
	jsonData, err := marshaler.MarshalTraces(td)
	if err != nil {
		log.Printf("Error marshaling traces: %v", err)
		return err
	}

	// Log the JSON representation of the traces
	log.Printf("Traces JSON: %s", string(jsonData))
	return nil
}
