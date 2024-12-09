// Package sdk provides real implementations of the Tracer and Span interfaces.
package sdk

import (
	"fmt"

	"github.com/ZiadMansourM/no-op/interfaces"
)

// RealSpan is a real implementation of Span.
type RealSpan struct {
	name string
}

// End finishes the span.
func (r *RealSpan) End() {
	fmt.Printf("Span ended: %s\n", r.name)
}

// RealTracer is a real implementation of Tracer.
type RealTracer struct{}

// StartSpan starts a real span.
func (r *RealTracer) StartSpan(name string) interfaces.Span {
	fmt.Printf("Span started: %s\n", name)
	return &RealSpan{name: name}
}
