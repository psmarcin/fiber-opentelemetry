package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lightstep/otel-launcher-go/launcher"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
)

func main() {
	// configuration
	// collector setup is out of the scope of this package so we don't focuse on it, just go to lightstep.com
	// and get credential
	launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("example-basic"),
		launcher.WithAccessToken("qpEupcxvxB4yfLqC0RLG1kf+2pa5afIeVhIR6mnE3pgSEV+K3V50/FyMGOHv3BhaL4c+haJg63Kr9Z4fa7adqczGCBFXlUHFcmybSMUu"),
	)

	// app setup
	app := fiber.New()

	// middleware use
	app.Use(fiberOtel.New(fiberOtel.Config{
		SpanName:     "http/request",
		LocalKeyName: "custom-local-key-to-store-otel-context",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// example how to get otel context form *fiber.Ctx
	app.Get("/nested/route", func(c *fiber.Ctx) error {
		// retrieve otel context from *fiber.Ctx
		ctx := fiberOtel.FromCtx(c)

		// use retrieved context
		_, span := fiberOtel.Tracer.Start(ctx, "nested-route-tracer")
		span.AddEvent("get-post")
		span.AddEvent("get-comments")
		span.AddEvent("get-author")
		defer span.End()

		return c.SendString("Additional trace has been sent")
	})

	app.Listen(":3000")
}
