# Fiber OpenTelemetry

OpenTelemetry trace middleware for [Fiber](https://github.com/gofiber/fiber) that adds traces to requests.

### Table of Contents

- [Signatures](#signatures)
- [Examples](#examples)
- [Config](#config)
- [Default Config](#default-config)

### Signatures

```go
func New(config ...Config) fiber.Handler
```

### Examples

Import the middleware package that is part of the Fiber web framework

```go
import (
    "github.com/gofiber/fiber/v2"
    "go.opentelemetry.io/otel/api/trace"
    "go.opentelemetry.io/otel/label"
    fiberOpentelemetry "github.com/psmarcin/fiber-opentelemetry"
)
```

After you initiate your Fiber app, you can use the following possibilities:

```go
// Default middleware config
tracer := trace.NoopTracerProvider().Tracer("test")
app.Use(fiberOpentelemetry.New(fiberOpentelemetry.Config{
    Tracer: trace,
}))

app.Get("/", func(ctx *fiber.Ctx) error {
    c := fiber_opentelemetry.FromCtx(ctx)
    
    ctx, span := tracer.Start(ctx, "trace-name")
    defer span.End()
    
    // attribute
    span.SetAttributes(label.String("attribute-name", "123"))
	
    // error
    span.RecordError(ctx, err)
    
    // event
    span.AddEvent(ctx, "event-name")
})
```

### Config

```go
// Config defines the config for middleware.
type Config struct {
    Tracer                trace.Tracer
    TracerStartAttributes []trace.SpanOption
    SpanName              string
    LocalKeyName          string
}
```

### Default Config

```go
var ConfigDefault = Config{
    SpanName:     "http/request",
    LocalKeyName: LocalsCtxKey,
    TracerStartAttributes: []trace.SpanOption{
        trace.WithSpanKind(trace.SpanKindServer),
        trace.WithNewRoot(),
        trace.WithRecord(),
    },
}
```
