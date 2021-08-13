package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func createSegment(c *fiber.Ctx, name string) func() {
	txn := FromContext(c)
	s := newrelic.Segment{
		Name:      "GET /users controller",
		StartTime: txn.StartSegmentNow(),
	}
	return s.End
}

var (
	sampleResponseBody = map[string]string{
		"id":   "17389123",
		"name": "ezekiel",
	}
	sampleErrorBody = map[string]string{
		"message": "wrong request",
		"code":    "red",
	}
)

func main() {
	app := fiber.New()
	nr, err := newrelic.NewApplication(newrelic.ConfigEnabled(true), newrelic.ConfigAppName("fiber-demo"), newrelic.ConfigLicense("license-key"))
	if err != nil {
		log.Fatal(err)
	}

	app.Use(Middleware(nr))

	app.Get("/users", func(c *fiber.Ctx) error {
		defer createSegment(c, "GET /users controller")()

		time.Sleep(10 * time.Millisecond)

		return c.Status(http.StatusOK).JSON(sampleResponseBody)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		defer createSegment(c, "POST /users controller")()

		time.Sleep(15 * time.Millisecond)
		c.Response().Header.Add("error", "true")
		return c.Status(http.StatusBadRequest).JSON(sampleErrorBody)
	})

	log.Fatal(app.Listen(":3000"))
}
