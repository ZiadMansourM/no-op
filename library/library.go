// Package library simulates an instrumented library using the OpenTelemetry API.
package library

import (
	"fmt"

	"github.com/ZiadMansourM/no-op/api"
)

// InstrumentedFunction demonstrates a function that uses the OpenTelemetry API.
func InstrumentedFunction() {
	tracer := api.Tracer()

	span := tracer.StartSpan("instrumented_function")
	fmt.Println("Doing some work in the instrumented function...")
	span.End()
}
