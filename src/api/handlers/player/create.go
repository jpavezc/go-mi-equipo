package player

import (
	"github.com/gin-gonic/gin"
	"mi-equipo-backend/src/api/internal/domain"
	"mi-equipo-backend/src/api/internal/services"
	"net/http"
)

func CreatePlayer(c *gin.Context) {
	var playerParams domain.Player
	if err := c.BindJSON(&playerParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	insertedId, err := player.CreatePlayerService(playerParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops!"})
	}

	c.JSON(200, gin.H{"insertedResult": insertedId})
}
