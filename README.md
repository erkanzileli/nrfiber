# nrfiber

Provides auto instrumentation for [NewRelic](https://newrelic.com) and [GoFiber](https://gofiber.io).

## Install

```shell
go get -u github.com/erkanzileli/nrfiber
```

## Usage

Register the middleware and use created transaction to add another segments.

```go
package main

import (
	"github.com/erkanzileli/nrfiber"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// Add the nrfiber middleware before other middlewares or routes
	app.Use(nrfiber.Middleware(app))

	// Use created transaction to create custom segments
	app.Get("/cart", func(ctx *fiber.Ctx) error {
		txn := FromContext(ctx)
		segment := txn.StartSegment("Price Calculation")
		defer segment.End()

		// calculate the price

		return nil
	})

}
```

### Contributing

Feel free to add anything useful or fix something.
