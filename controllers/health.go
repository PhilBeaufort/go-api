package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Ping
// @Description  Responds with "pong"
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// @Summary      Health Check
// @Description  Returns 200 if server is healthy
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
