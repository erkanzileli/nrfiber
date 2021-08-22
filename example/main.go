package main

import (
	"fmt"
	"github.com/erkanzileli/nrfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
)

type customErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ce customErr) Error() string {
	return fmt.Sprintf("code: %d, message: %s", ce.Code, ce.Message)
}

func main() {
	app := fiber.New()
	nr, err := newrelic.NewApplication(newrelic.ConfigEnabled(true), newrelic.ConfigAppName("demo"), newrelic.ConfigLicense("license-key"))
	if err != nil {
		log.Fatal(err)
	}
	app.Use(nrfiber.Middleware(nr, nrfiber.ConfigNoticeErrorEnabled(true)))
	app.Get("/give-me-error", func(ctx *fiber.Ctx) error {
		err := customErr{Message: "wrong request", Code: 4329}
		ctx.Status(http.StatusBadRequest).JSON(err)
		return err
	})
	app.Listen(":3000")
}
