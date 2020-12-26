package fiber_otel

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/semconv"
)

const LocalsCtxKey = "otel-ctx"

var Tracer = otel.Tracer("fiber-otel-router")

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// concat all span options, dynamic and static
		spanOptions := concatSpanOptions(
			[]trace.SpanOption{
				trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
				trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
			},
			cfg.TracerStartAttributes,
		)

		ctx, span := Tracer.Start(
			c.Context(),
			cfg.SpanName,
			spanOptions...,
		)

		c.Locals(LocalsCtxKey, ctx)
		defer span.End()
		err := c.Next()

		span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(c.Response().StatusCode())...)

		return err
	}
}


func concatSpanOptions(sources ...[]trace.SpanOption) []trace.SpanOption {
	var spanOptions []trace.SpanOption
	for _, source := range sources {
		for _, option := range source {
			spanOptions = append(spanOptions, option)
		}
	}

	return spanOptions
}