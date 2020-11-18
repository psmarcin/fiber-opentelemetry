package fiber_opentelemetry

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/api/trace"
)

func FromCtx(ctx *fiber.Ctx) context.Context {
	otelCtx := ctx.Locals(LocalsCtxKey).(context.Context)
	return otelCtx
}

func SpanFromCtx(ctx *fiber.Ctx) trace.Span {
	otelCtx := FromCtx(ctx)
	if otelCtx == nil {
		return trace.SpanFromContext(ctx.Context())
	}

	return trace.SpanFromContext(otelCtx)
}
