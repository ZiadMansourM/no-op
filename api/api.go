// Package api provides the global TracerProvider and no-op implementations
// for the OpenTelemetry API. It ensures instrumented code works even if
// the SDK is uninitialized.
package api

import "github.com/ZiadMansourM/no-op/interfaces"

// noOpSpan is a no-op implementation of interfaces.Span.
type noOpSpan struct{}

// End is a no-op method.
func (n *noOpSpan) End() {}

// noOpTracer is a no-op implementation of interfaces.Tracer.
type noOpTracer struct{}

// StartSpan returns a no-op span.
func (n *noOpTracer) StartSpan(name string) interfaces.Span {
	return &noOpSpan{}
}

// Global TracerProvider
var tracerProvider interfaces.Tracer = &noOpTracer{}

// Tracer provides the current tracer.
func Tracer() interfaces.Tracer {
	return tracerProvider
}

// SetTracerProvider sets the tracer provider.
func SetTracerProvider(provider interfaces.Tracer) {
	tracerProvider = provider
}
