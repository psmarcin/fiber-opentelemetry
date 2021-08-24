package fiber_otel

import (
	"go.opentelemetry.io/otel/trace"
)

// Config defines the config for middleware.
type Config struct {
	Tracer                trace.Tracer
	TracerStartAttributes []trace.SpanStartOption
	// SpanName is a template for span naming.
	// The scope is fiber context.
	SpanName     string
	LocalKeyName string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	SpanName:     "{{ .Method }} {{ .Route.Path }}",
	LocalKeyName: LocalsCtxKey,
	TracerStartAttributes: []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithNewRoot(),
	},
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.SpanName == "" {
		cfg.SpanName = ConfigDefault.SpanName
	}

	if cfg.LocalKeyName == "" {
		cfg.LocalKeyName = ConfigDefault.LocalKeyName
	}

	if len(cfg.TracerStartAttributes) == 0 {
		cfg.TracerStartAttributes = ConfigDefault.TracerStartAttributes
	}

	return cfg
}
