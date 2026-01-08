package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	services "github.com/0xRichardL/otel-prom-practice/game/internal/services"
)

type App struct {
	dice     *services.Dice
	roulette *services.Roulette
	router   *gin.Engine
}

func NewApp(dice *services.Dice, roulette *services.Roulette) *App {
	return &App{
		dice:     dice,
		roulette: roulette,
	}
}

// SetupRouter configures all HTTP routes and middleware
func (s *App) SetupRouter() *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(s.loggingMiddleware())

	// Routes
	router.GET("/", s.handleRoot)
	router.GET("/health", s.handleHealth)
	router.GET("/dice/roll", s.handleDiceRoll)
	router.GET("/roulette/spin", s.handleRouletteSpin)

	s.router = router
	return router
}

// loggingMiddleware logs request details
func (s *App) loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		duration := time.Since(start)
		log.Printf("[%s] %s %s?%s - Status: %d - Duration: %v",
			c.Request.Method, path, path, query, c.Writer.Status(), duration)
	}
}

func (s *App) handleRoot(c *gin.Context) {
	response := gin.H{
		"service":   "game-service",
		"version":   "1.0.0",
		"timestamp": time.Now().Unix(),
		"endpoints": []string{
			"/health",
			"/dice/roll?bet=<amount>&type=<single|odd|even|high|low>&value=<1-6>",
			"/roulette/spin?bet=<amount>&type=<single|red|black|odd|even|dozen|column>&value=<0-36>",
		},
	}

	c.JSON(http.StatusOK, response)
}

func (s *App) handleHealth(c *gin.Context) {
	response := gin.H{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
