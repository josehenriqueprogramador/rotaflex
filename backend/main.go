package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/rides/available", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"rides": []string{},
			})
		})

		api.POST("/rides", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "corrida criada com sucesso",
			})
		})

		api.PATCH("/rides/:id/accept", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "corrida aceita",
			})
		})
	}

	r.Run(":8000")
}
