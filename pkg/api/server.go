package api

import (
	"bordeaux-matching-engine-exercise/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine model.MatchingEngine
	router *gin.Engine
}

func NewServer() (*Server, error) {
	engine := model.NewMatchingEngine()

	server := &Server{engine: *engine}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// enable CORS
	router.Use(CORSMiddleware())

	router.POST("/order", func(c *gin.Context) {
		var order model.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newOrder := server.engine.PlaceOrder(order.OrderType, order.Side, order.Price, order.Quantity)
		c.JSON(http.StatusOK, newOrder)
	})

	router.GET("/orderbook", func(c *gin.Context) {
		orderBook := server.engine.GetOrderBook()
		c.JSON(http.StatusOK, orderBook)
	})

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

// CORSMiddleware sets the CORS headers to allow all origins
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
