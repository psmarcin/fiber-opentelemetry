module github.com/psmarcin/fiber-opentelemetry/example/trace

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofiber/fiber/v2 v2.15.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/psmarcin/fiber-opentelemetry v0.5.1
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/psmarcin/fiber-opentelemetry => ./../../
