package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/0xRichardL/otel-prom-practice/game/internal/errors"
	"github.com/gin-gonic/gin"
)

// DiceRollRequest represents a dice roll request
type DiceRollRequest struct {
	Bet      float64 `form:"bet" binding:"required,gt=0"`
	BetType  string  `form:"type" binding:"required,oneof=single odd even high low"`
	BetValue int     `form:"value" binding:"omitempty,min=1,max=6"`
}

func (s *App) handleDiceRoll(c *gin.Context) {
	start := time.Now()
	log.Printf("[DICE] Request started: method=%s path=%s", c.Request.Method, c.Request.URL.Path)

	var req DiceRollRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errors.RespondError(c, err)
		return
	}

	// Play the game
	result, err := s.dice.Roll(req.Bet, req.BetType, req.BetValue)
	if err != nil {
		log.Printf("[DICE] Error: %v", err)
		errors.RespondError(c, err)
		return
	}

	duration := time.Since(start)
	log.Printf("[DICE] Request completed: bet=%.2f type=%s result=%d won=%t payout=%.2f duration=%v",
		req.Bet, req.BetType, result.Roll, result.Won, result.Payout, duration)

	// Record the dice roll metric
	s.metrics.RecordDiceRoll(c.Request.Context())

	c.JSON(http.StatusOK, result)
}
