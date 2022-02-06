package nrfiber

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/valyala/fasthttp"
)

// FromContext returns the Transaction from the context if present, and nil
// otherwise.
func FromContext(c *fiber.Ctx) *newrelic.Transaction {
	return newrelic.FromContext(c.UserContext())
}

// Middleware creates Fiber middleware that instruments requests.
//
//	app := fiber.New()
//	// Add the nrfiber middleware before other middlewares or routes:
//	app.Use(nrfiber.Middleware(app))
func Middleware(app *newrelic.Application, configs ...*config) fiber.Handler {
	if nil == app {
		return func(c *fiber.Ctx) error {
			return nil
		}
	}

	configMap := createConfigMap(configs...)
	noticeErrorEnabled := noticeErrorEnabled(configMap)

	return func(c *fiber.Ctx) error {
		txn := app.StartTransaction(createTransactionName(c))
		defer txn.End()

		txn.SetWebRequestHTTP(createHttpRequest(c))

		c.SetUserContext(newrelic.NewContext(c.UserContext(), txn))

		err := c.Next()
		statusCode := c.Context().Response.StatusCode()

		if err != nil {
			if fiberErr, ok := err.(*fiber.Error); ok {
				statusCode = fiberErr.Code
			}
			if noticeErrorEnabled {
				txn.NoticeError(err)
			}
		}

		txn.SetWebResponse(nil).WriteHeader(statusCode)
		return nil
	}
}

func createTransactionName(c *fiber.Ctx) string {
	return fmt.Sprintf("%s %s", c.Request().Header.Method(), c.Request().URI().Path())
}

func convertRequestHeaders(fastHttpHeaders *fasthttp.RequestHeader) http.Header {
	headers := make(http.Header)

	fastHttpHeaders.VisitAll(func(k, v []byte) {
		headers.Set(string(k), string(v))
	})

	return headers
}

func createHttpRequest(c *fiber.Ctx) *http.Request {
	reqHeaders := convertRequestHeaders(&c.Request().Header)

	reqHost := reqHeaders.Get("Host")
	if reqHost == "" {
		reqHost = string(c.Request().URI().Host())
	}

	return &http.Request{
		Method: c.Method(),
		URL: &url.URL{
			Scheme:   string(c.Request().URI().Scheme()),
			Host:     reqHost,
			Path:     string(c.Request().URI().Path()),
			RawQuery: string(c.Request().URI().QueryString()),
		},
		Header: reqHeaders,
		Host:   reqHost,
	}
}
