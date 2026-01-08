package services

import (
	"fmt"
	"math/rand"
	"time"
)

type Roulette struct {
}

func NewRoulette() *Roulette {
	return &Roulette{}
}

// RouletteResult represents the result of a roulette spin
type RouletteResult struct {
	Number     int     `json:"number"`
	Color      string  `json:"color"`
	Bet        float64 `json:"bet"`
	BetType    string  `json:"bet_type"`
	BetValue   int     `json:"bet_value,omitempty"`
	Won        bool    `json:"won"`
	Payout     float64 `json:"payout"`
	Multiplier float64 `json:"multiplier"`
	Timestamp  int64   `json:"timestamp"`
}

// Roulette wheel configuration (European-style with single zero)
var (
	redNumbers   = []int{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
	blackNumbers = []int{2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35}
)

// SpinRoulette simulates a roulette game with various bet types
func (r *Roulette) Spin(bet float64, betType string, betValue int) (*RouletteResult, error) {
	// Simulate wheel spinning time (roulette is slower than dice)
	time.Sleep(time.Millisecond * time.Duration(50+rand.Intn(100)))

	if bet <= 0 {
		return nil, fmt.Errorf("bet must be greater than 0")
	}

	// Spin the wheel (0-36)
	number := rand.Intn(37)

	// Determine color
	color := getColor(number)

	result := &RouletteResult{
		Number:    number,
		Color:     color,
		Bet:       bet,
		BetType:   betType,
		BetValue:  betValue,
		Timestamp: time.Now().Unix(),
	}

	// Determine if the player won based on bet type
	switch betType {
	case "single":
		if betValue < 0 || betValue > 36 {
			return nil, fmt.Errorf("for 'single' bet, value must be between 0 and 36")
		}
		result.Won = number == betValue
		result.Multiplier = 35.0 // 35:1 payout for single number

	case "red":
		result.Won = color == "red"
		result.Multiplier = 2.0 // 1:1 payout

	case "black":
		result.Won = color == "black"
		result.Multiplier = 2.0

	case "odd":
		result.Won = number > 0 && number%2 == 1
		result.Multiplier = 2.0

	case "even":
		result.Won = number > 0 && number%2 == 0
		result.Multiplier = 2.0

	case "low":
		// 1-18
		result.Won = number >= 1 && number <= 18
		result.Multiplier = 2.0

	case "high":
		// 19-36
		result.Won = number >= 19 && number <= 36
		result.Multiplier = 2.0

	case "dozen1":
		// 1-12
		result.Won = number >= 1 && number <= 12
		result.Multiplier = 3.0 // 2:1 payout

	case "dozen2":
		// 13-24
		result.Won = number >= 13 && number <= 24
		result.Multiplier = 3.0

	case "dozen3":
		// 25-36
		result.Won = number >= 25 && number <= 36
		result.Multiplier = 3.0

	case "column1":
		// 1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34
		result.Won = number > 0 && number%3 == 1
		result.Multiplier = 3.0

	case "column2":
		// 2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35
		result.Won = number > 0 && number%3 == 2
		result.Multiplier = 3.0

	case "column3":
		// 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36
		result.Won = number > 0 && number%3 == 0
		result.Multiplier = 3.0

	default:
		return nil, fmt.Errorf("invalid bet type: %s (valid types: single, red, black, odd, even, low, high, dozen1-3, column1-3)", betType)
	}

	// Calculate payout
	if result.Won {
		result.Payout = bet * result.Multiplier
	} else {
		result.Payout = 0
	}

	return result, nil
}

// getColor returns the color of a roulette number
func getColor(number int) string {
	if number == 0 {
		return "green"
	}

	for _, n := range redNumbers {
		if n == number {
			return "red"
		}
	}

	for _, n := range blackNumbers {
		if n == number {
			return "black"
		}
	}

	return "unknown"
}
