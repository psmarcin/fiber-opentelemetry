module github.com/psmarcin/fiber-opentelemetry/example/trace

go 1.15

require (
	github.com/gofiber/fiber/v2 v2.7.1
	github.com/lightstep/otel-launcher-go v0.18.0
	github.com/psmarcin/fiber-opentelemetry v0.4.0
)

replace github.com/psmarcin/fiber-opentelemetry => ./../../
