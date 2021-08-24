package main

import (
	"github.com/gofiber/fiber/v2"
	fiberOtel "github.com/psmarcin/fiber-opentelemetry/pkg/fiber-otel"
)

func main() {
	// app setup
	app := fiber.New()

	// middleware use
	app.Use(fiberOtel.New(fiberOtel.Config{
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
