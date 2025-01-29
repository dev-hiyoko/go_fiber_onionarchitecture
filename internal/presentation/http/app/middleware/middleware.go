package middleware

import (
	"strings"
	"time"

	"hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const (
	HeaderNameCache     = "X-Cache"
	CacheExp            = 30 * time.Minute
	RequestIDContextKey = "requestid"
	RequestTimeoutTime  = 30 * time.Second
)

func NewMiddleware(app *fiber.App) {
	// recover
	app.Use(recover.New())

	// caching
	// app.Use(cache.New(cache.Config{
	//	Next: func(c *fiber.Ctx) bool {
	//		return c.Query("noCache") == "true"
	//	},
	//	Expiration:   CacheExp,
	//	CacheHeader:  HeaderNameCache,
	//	CacheControl: true,
	// }))

	// etag
	// app.Use(etag.New(etag.Config{
	//	Weak: true,
	// }))

	// encoding
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: utils.Env("FRONT_APP_URL").GetString("*"),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: strings.Join([]string{
			"Authorization",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
		}, ","),
	}))

	// early data
	app.Use(earlydata.New(earlydata.Config{
		Error: fiber.ErrTooEarly,
	}))

	// helmet
	app.Use(helmet.New())

	// idempotency
	app.Use(idempotency.New())

	// limiter
	app.Use(limiter.New(limiter.Config{
		Max:               utils.Env("LIMITER_MAX").GetInt(20),
		Expiration:        utils.Env("LIMITER_EXPIRATION").GetDuration(30 * time.Second),
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// request id
	app.Use(requestid.New(requestid.Config{
		Header:     fiber.HeaderXRequestID,
		ContextKey: RequestIDContextKey,
	}))

	// logger
	app.Use(func(c *fiber.Ctx) error {
		log.With("request id", c.Locals(RequestIDContextKey))
		return c.Next()
	})

	// access log
	app.Use(accessLogger())

	// timeout
	app.Use(timeoutMiddleware(RequestTimeoutTime))
}
