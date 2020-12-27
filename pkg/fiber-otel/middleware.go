package fiber_otel

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
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
				trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
				trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
				trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
				trace.WithAttributes(semconv.NetHostIPKey.String(c.IP())),
				trace.WithAttributes(semconv.HTTPUserAgentKey.String(string(c.Request().Header.UserAgent()))),
				trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
				trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
				trace.WithAttributes(semconv.NetTransportTCP),
				trace.WithSpanKind(trace.SpanKindServer),
				// TODO:
				// - x-forwarded-for
				// -
			},
			cfg.TracerStartAttributes,
		)

		otelCtx, span := Tracer.Start(
			c.Context(),
			c.Route().Path,
			spanOptions...,
		)

		c.Locals(LocalsCtxKey, otelCtx)
		defer span.End()

		err := c.Next()

		statusCode := c.Response().StatusCode()
		attrs := semconv.HTTPAttributesFromHTTPStatusCode(statusCode)
		spanStatus, spanMessage := semconv.SpanStatusFromHTTPStatusCode(statusCode)
		span.SetAttributes(attrs...)
		span.SetStatus(spanStatus, spanMessage)

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
