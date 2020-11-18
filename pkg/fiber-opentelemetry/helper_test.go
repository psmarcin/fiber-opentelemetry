package fiber_opentelemetry

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/api/trace"
	"net/http/httptest"
	"testing"
)

func TestFromCtx(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{
		Tracer: trace.NoopTracerProvider().Tracer("test"),
	}))

	var ctxVal context.Context

	app.Use(func(c *fiber.Ctx) error {
		ctxVal = FromCtx(c)
		return c.Next()
	})

	_, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.NoError(t, err)
	assert.NotEqual(t, ctxVal, nil)
}

func TestSpanFromCtx(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{
		Tracer: trace.NoopTracerProvider().Tracer("test"),
	}))

	var ctxVal trace.Span

	app.Use(func(c *fiber.Ctx) error {
		ctxVal = SpanFromCtx(c)
		return c.Next()
	})

	_, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.NoError(t, err)
	assert.NotEqual(t, ctxVal, nil)
}
