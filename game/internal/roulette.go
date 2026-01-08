package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/0xRichardL/otel-prom-practice/game/internal/errors"
	"github.com/gin-gonic/gin"
)

// RouletteSpinRequest represents a roulette spin request
type RouletteSpinRequest struct {
	Bet      float64 `form:"bet" binding:"required,gt=0"`
	BetType  string  `form:"type" binding:"required,oneof=single red black odd even low high dozen1 dozen2 dozen3 column1 column2 column3"`
	BetValue int     `form:"value" binding:"omitempty,min=0,max=36"`
}

func (s *App) handleRouletteSpin(c *gin.Context) {
	start := time.Now()
	log.Printf("[ROULETTE] Request started: method=%s path=%s", c.Request.Method, c.Request.URL.Path)

	var req RouletteSpinRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errors.RespondError(c, err)
		return
	}

	// Use -1 as default if value not provided
	value := req.BetValue
	if value == 0 && req.BetType != "single" {
		value = -1
	}

	// Play the game
	result, err := s.roulette.Spin(req.Bet, req.BetType, value)
	if err != nil {
		log.Printf("[ROULETTE] Error: %v", err)
		errors.RespondError(c, err)
		return
	}

	duration := time.Since(start)
	log.Printf("[ROULETTE] Request completed: bet=%.2f type=%s result=%d color=%s won=%t payout=%.2f duration=%v",
		req.Bet, req.BetType, result.Number, result.Color, result.Won, result.Payout, duration)

	c.JSON(http.StatusOK, result)
}
