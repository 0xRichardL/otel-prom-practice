package utils

import (
	"math/rand/v2"

	"github.com/0xRichardL/otel-prom-practice/game/internal/errors"
	"github.com/gin-gonic/gin"
)

// UnstableOperation Randomly return error based on the given rate
func UnstableOperation(rate float64) error {
	if rand.Float64() < rate {
		return errors.NewApplicationError(
			"random unstable error occurred",
			errors.AppErrorTypeInternal,
		)
	}
	return nil
}

func UnstableMiddleware(rate float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := UnstableOperation(rate); err != nil {
			errors.RespondError(c, err)
			return
		}
		c.Next()
	}
}
