package metrics

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	MeterName = "game-service"
)

// AppMetrics holds our metrics
type AppMetrics struct {
	diceRollsTotal metric.Int64Counter
}

// NewAppMetrics creates a simple counter metric
func NewAppMetrics() (*AppMetrics, error) {
	meter := otel.Meter(MeterName)

	// Create a counter to track total dice rolls
	diceRollsTotal, err := meter.Int64Counter(
		"dice_rolls_total",
		metric.WithDescription("Total number of dice rolls"),
		metric.WithUnit("{roll}"),
	)
	if err != nil {
		return nil, err
	}

	return &AppMetrics{
		diceRollsTotal: diceRollsTotal,
	}, nil
}
