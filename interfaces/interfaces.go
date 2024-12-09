// Package interfaces defines the core interfaces for Tracer and Span
package interfaces

// Span defines the interface for a span.
type Span interface {
	End()
}

// Tracer defines the interface for a tracer.
type Tracer interface {
	StartSpan(name string) Span
}
