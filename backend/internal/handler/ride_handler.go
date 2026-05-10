package handler

import (
	"github.com/gin-gonic/gin"
	"rotasflex/internal/service"
)

type RideHandler struct {
	service *service.RideService
}

func NewRideHandler(s *service.RideService) *RideHandler {
	return &RideHandler{service: s}
}

func (h *RideHandler) Create(c *gin.Context) {
	ride := h.service.CreateRide(1)

	c.JSON(201, gin.H{
		"ride": ride,
	})
}
