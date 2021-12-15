package otelsql

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type traceAttributes []attribute.KeyValue

// Config is used to configure the go-restful middleware.
type config struct {
	traceProvider   trace.TracerProvider
	traceAttributes traceAttributes
}

// Option specifies instrumentation configuration options.
type Option func(*config)

// WithTracer configures the interceptor with the specified trace provider.
func WithTraceProvider(traceProvider trace.TracerProvider) Option {
	return func(cfg *config) {
		cfg.traceProvider = traceProvider
	}
}

// WithTracer configures the interceptor to attach the default KeyValues.
func WithTraceAttributes(traceAttributes []attribute.KeyValue) Option {
	return func(cfg *config) {
		cfg.traceAttributes = traceAttributes
	}
}
