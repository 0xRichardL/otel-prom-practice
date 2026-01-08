package services

import (
	"fmt"
	"math/rand"
	"time"
)

type Dice struct {
}

func NewDice() *Dice {
	return &Dice{}
}

// DiceResult represents the result of a dice roll
type DiceResult struct {
	Roll       int     `json:"roll"`
	Bet        float64 `json:"bet"`
	BetType    string  `json:"bet_type"`
	BetValue   int     `json:"bet_value,omitempty"`
	Won        bool    `json:"won"`
	Payout     float64 `json:"payout"`
	Multiplier float64 `json:"multiplier"`
	Timestamp  int64   `json:"timestamp"`
}

// Roll simulates a dice roll game with various bet types
func (*Dice) Roll(bet float64, betType string, betValue int) (*DiceResult, error) {
	// Simulate some processing time
	time.Sleep(time.Millisecond * time.Duration(10+rand.Intn(40)))

	if bet <= 0 {
		return nil, fmt.Errorf("bet must be greater than 0")
	}

	// Roll the dice (1-6)
	roll := rand.Intn(6) + 1

	result := &DiceResult{
		Roll:      roll,
		Bet:       bet,
		BetType:   betType,
		BetValue:  betValue,
		Timestamp: time.Now().Unix(),
	}

	// Determine if the player won based on bet type
	switch betType {
	case "single":
		if betValue < 1 || betValue > 6 {
			return nil, fmt.Errorf("for 'single' bet, value must be between 1 and 6")
		}
		result.Won = roll == betValue
		result.Multiplier = 6.0 // 6x payout for single number

	case "odd":
		result.Won = roll%2 == 1
		result.Multiplier = 2.0 // 2x payout for odd/even

	case "even":
		result.Won = roll%2 == 0
		result.Multiplier = 2.0

	case "high":
		// High: 4, 5, 6
		result.Won = roll >= 4
		result.Multiplier = 2.0

	case "low":
		// Low: 1, 2, 3
		result.Won = roll <= 3
		result.Multiplier = 2.0

	default:
		return nil, fmt.Errorf("invalid bet type: %s (valid types: single, odd, even, high, low)", betType)
	}

	// Calculate payout
	if result.Won {
		result.Payout = bet * result.Multiplier
	} else {
		result.Payout = 0
	}

	return result, nil
}
