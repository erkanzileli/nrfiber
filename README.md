# nrfiber

Provides auto instrumentation for [NewRelic](https://newrelic.com) and [GoFiber](https://gofiber.io).

## Install

```shell
go get -u github.com/erkanzileli/nrfiber
```

## Usage

Register the middleware and use created transaction to add another segments. Basic usage is below

```go
package main

import (
	"github.com/erkanzileli/nrfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
)

func main() {
	app := fiber.New()
	nr, err := newrelic.NewApplication(newrelic.ConfigEnabled(true), newrelic.ConfigAppName("demo"), newrelic.ConfigLicense("license-key"))
	if err != nil {
		log.Fatal(err)
	}

	// Add the nrfiber middleware before other middlewares or routes
	app.Use(nrfiber.Middleware(nr))

	// Use created transaction to create custom segments
	app.Get("/cart", func(ctx *fiber.Ctx) error {
		txn := nrfiber.FromContext(ctx)
		segment := txn.StartSegment("Price Calculation")
		defer segment.End()

		// calculate the price

		return nil
	})
	app.Listen(":3000")
}
```

## Guides

- [Notice Custom Errors](docs/notice-custom-errors.md)

### Contributing

Feel free to add anything useful or fix something.
