package v1

import (
	"net/http"

	"backend/internal/services"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func MatchPlayers(c *gin.Context) {
	userID, _ := c.Get("userID")

	roomID, err := services.MatchPlayers(userID.(uint))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"room_id": roomID})
}
