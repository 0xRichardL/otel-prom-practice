# Game Service API

A simple HTTP game service with Dice and Roulette games, designed for OpenTelemetry practice.

## 🚀 Running the Service

```bash
# Run the service
go run cmd/main.go

# Or with custom port
PORT=8080 go run cmd/main.go
```

## 📋 API Endpoints

### Health Check

```bash
# Check service health
curl http://localhost:8080/health
```

### Root Endpoint

```bash
# Get service info
curl http://localhost:8080/
```

---

## 🎲 Dice Game

Roll a six-sided die with various betting options.

### Single Number Bet (6x payout)

```bash
# Bet $10 on rolling a 5
curl "http://localhost:8080/dice/roll?bet=10&type=single&value=5"

# Bet $25 on rolling a 3
curl "http://localhost:8080/dice/roll?bet=25&type=single&value=3"

# Bet $100 on rolling a 1
curl "http://localhost:8080/dice/roll?bet=100&type=single&value=1"
```

### Odd/Even Bet (2x payout)

```bash
# Bet $20 on odd numbers (1, 3, 5)
curl "http://localhost:8080/dice/roll?bet=20&type=odd"

# Bet $50 on even numbers (2, 4, 6)
curl "http://localhost:8080/dice/roll?bet=50&type=even"
```

### High/Low Bet (2x payout)

```bash
# Bet $15 on high numbers (4, 5, 6)
curl "http://localhost:8080/dice/roll?bet=15&type=high"

# Bet $30 on low numbers (1, 2, 3)
curl "http://localhost:8080/dice/roll?bet=30&type=low"
```

### Error Cases

```bash
# Invalid bet type
curl "http://localhost:8080/dice/roll?bet=10&type=invalid"

# Missing bet amount
curl "http://localhost:8080/dice/roll?type=odd"

# Invalid single number value
curl "http://localhost:8080/dice/roll?bet=10&type=single&value=7"

# Negative bet
curl "http://localhost:8080/dice/roll?bet=-10&type=odd"
```

---

## 🎰 Roulette Game

Spin a roulette wheel (0-36) with various betting options.

### Single Number Bet (35x payout)

```bash
# Bet $10 on number 17 (lucky number!)
curl "http://localhost:8080/roulette/spin?bet=10&type=single&value=17"

# Bet $5 on number 0 (green)
curl "http://localhost:8080/roulette/spin?bet=5&type=single&value=0"

# Bet $20 on number 7
curl "http://localhost:8080/roulette/spin?bet=20&type=single&value=7"

# Bet $50 on number 23
curl "http://localhost:8080/roulette/spin?bet=50&type=single&value=23"
```

### Color Bets (2x payout)

```bash
# Bet $25 on red
curl "http://localhost:8080/roulette/spin?bet=25&type=red"

# Bet $30 on black
curl "http://localhost:8080/roulette/spin?bet=30&type=black"
```

### Odd/Even Bet (2x payout)

```bash
# Bet $40 on odd numbers
curl "http://localhost:8080/roulette/spin?bet=40&type=odd"

# Bet $35 on even numbers
curl "http://localhost:8080/roulette/spin?bet=35&type=even"
```

### High/Low Bet (2x payout)

```bash
# Bet $20 on low numbers (1-18)
curl "http://localhost:8080/roulette/spin?bet=20&type=low"

# Bet $45 on high numbers (19-36)
curl "http://localhost:8080/roulette/spin?bet=45&type=high"
```

### Dozen Bets (3x payout)

```bash
# Bet $15 on first dozen (1-12)
curl "http://localhost:8080/roulette/spin?bet=15&type=dozen1"

# Bet $20 on second dozen (13-24)
curl "http://localhost:8080/roulette/spin?bet=20&type=dozen2"

# Bet $25 on third dozen (25-36)
curl "http://localhost:8080/roulette/spin?bet=25&type=dozen3"
```

### Column Bets (3x payout)

```bash
# Bet $10 on column 1 (1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34)
curl "http://localhost:8080/roulette/spin?bet=10&type=column1"

# Bet $15 on column 2 (2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35)
curl "http://localhost:8080/roulette/spin?bet=15&type=column2"

# Bet $20 on column 3 (3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36)
curl "http://localhost:8080/roulette/spin?bet=20&type=column3"
```

### Error Cases

```bash
# Invalid bet type
curl "http://localhost:8080/roulette/spin?bet=10&type=invalid"

# Missing bet amount
curl "http://localhost:8080/roulette/spin?type=red"

# Invalid single number value
curl "http://localhost:8080/roulette/spin?bet=10&type=single&value=37"

# Negative bet
curl "http://localhost:8080/roulette/spin?bet=-10&type=red"
```

---

## 📊 Response Format

### Successful Dice Response

```json
{
  "roll": 5,
  "bet": 10,
  "bet_type": "single",
  "bet_value": 5,
  "won": true,
  "payout": 60,
  "multiplier": 6,
  "timestamp": 1704672000
}
```

### Successful Roulette Response

```json
{
  "number": 17,
  "color": "black",
  "bet": 10,
  "bet_type": "single",
  "bet_value": 17,
  "won": true,
  "payout": 350,
  "multiplier": 35,
  "timestamp": 1704672000
}
```

### Error Response

```json
{
  "error": "Bad Request",
  "code": 400,
  "message": "Invalid bet amount"
}
```

---

## 🎯 Bet Types Summary

### Dice Game

| Bet Type | Description          | Payout Multiplier | Example                       |
| -------- | -------------------- | ----------------- | ----------------------------- |
| `single` | Exact number (1-6)   | 6x                | `?bet=10&type=single&value=5` |
| `odd`    | Odd numbers (1,3,5)  | 2x                | `?bet=10&type=odd`            |
| `even`   | Even numbers (2,4,6) | 2x                | `?bet=10&type=even`           |
| `high`   | High numbers (4-6)   | 2x                | `?bet=10&type=high`           |
| `low`    | Low numbers (1-3)    | 2x                | `?bet=10&type=low`            |

### Roulette Game

| Bet Type  | Description         | Payout Multiplier | Example                        |
| --------- | ------------------- | ----------------- | ------------------------------ |
| `single`  | Exact number (0-36) | 35x               | `?bet=10&type=single&value=17` |
| `red`     | Red numbers         | 2x                | `?bet=10&type=red`             |
| `black`   | Black numbers       | 2x                | `?bet=10&type=black`           |
| `odd`     | Odd numbers (1-35)  | 2x                | `?bet=10&type=odd`             |
| `even`    | Even numbers (2-36) | 2x                | `?bet=10&type=even`            |
| `low`     | Numbers 1-18        | 2x                | `?bet=10&type=low`             |
| `high`    | Numbers 19-36       | 2x                | `?bet=10&type=high`            |
| `dozen1`  | Numbers 1-12        | 3x                | `?bet=10&type=dozen1`          |
| `dozen2`  | Numbers 13-24       | 3x                | `?bet=10&type=dozen2`          |
| `dozen3`  | Numbers 25-36       | 3x                | `?bet=10&type=dozen3`          |
| `column1` | Column 1 (1,4,7...) | 3x                | `?bet=10&type=column1`         |
| `column2` | Column 2 (2,5,8...) | 3x                | `?bet=10&type=column2`         |
| `column3` | Column 3 (3,6,9...) | 3x                | `?bet=10&type=column3`         |

---

## 📥 Importing to Postman

1. Copy any curl command from above
2. In Postman, click **Import** > **Raw text**
3. Paste the curl command
4. Click **Continue** > **Import**
5. The request will be added to your collection

Alternatively, save all commands to a file and import as a collection:

```bash
# Save this README as a reference
# Or create a Postman collection JSON from these curl commands
```

---

## 🔧 Development

```bash
# Install dependencies
go mod download

# Run tests (if implemented)
go test ./...

# Build binary
go build -o game-service cmd/main.go

# Run binary
./game-service
```

---

## 📝 Notes for OpenTelemetry Integration

This service is designed with observability in mind:

- **Variable latency**: Dice (10-50ms), Roulette (50-150ms)
- **Multiple code paths**: Different bet types create diverse traces
- **Structured logging**: Ready for log correlation
- **Error scenarios**: Test error handling and alerting
- **Business metrics**: Win rates, bet amounts, game popularity

Perfect for learning:

- Distributed tracing
- Metrics collection
- Log aggregation
- Service monitoring
