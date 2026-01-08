package metrics

import "context"

// RecordDiceRoll increments the dice roll counter by 1
func (m *AppMetrics) RecordDiceRoll(ctx context.Context) {
	m.diceRollsTotal.Add(ctx, 1)
}
