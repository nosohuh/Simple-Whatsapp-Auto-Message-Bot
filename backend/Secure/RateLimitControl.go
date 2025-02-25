package securecontroller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(maxRequests int) fiber.Handler {
	config := limiter.Config{
		Max:               maxRequests,
		Expiration:        5 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).
				JSON(fiber.Map{
					"error":   "Rate limit exceeded",
					"status":  fiber.StatusTooManyRequests,
					"success": false,
				})
		},
	}

	return limiter.New(config)
}
