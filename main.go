// Package main demonstrates the usage of the OpenTelemetry API and SDK.
package main

import (
	"fmt"

	"github.com/ZiadMansourM/no-op/api"
	"github.com/ZiadMansourM/no-op/library"
	"github.com/ZiadMansourM/no-op/sdk"
)

func main() {
	fmt.Println("Running without SDK initialized:")
	library.InstrumentedFunction()

	fmt.Println("\nInitializing SDK...")
	api.SetTracerProvider(&sdk.RealTracer{})
	library.InstrumentedFunction()
}
