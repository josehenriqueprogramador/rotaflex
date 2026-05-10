package handler

import (
	"github.com/gin-gonic/gin"
	"rotasflex/internal/service"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	s := &Server{
		router: r,
	}

	s.routes()
	return s
}

func (s *Server) routes() {

	// 🔥 service (regra de negócio)
	rideService := service.NewRideService()

	// 🔌 handler (HTTP)
	rideHandler := NewRideHandler(rideService)

	api := s.router.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.POST("/rides", rideHandler.Create)
}

func (s *Server) Run() {
	s.router.Run(":8000")
}
