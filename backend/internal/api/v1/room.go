package v1

import (
	"net/http"
	"strconv"

	"backend/internal/models"
	"backend/internal/services"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

type CreateRoomRequest struct {
	Name      string `json:"name" binding:"required"`
	IsPrivate bool   `json:"is_private"`
	PassCode  string `json:"PassCode"`
}

func CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")

	room := models.Room{
		Name:      req.Name,
		IsPrivate: req.IsPrivate,
		PassCode:  req.PassCode,
	}

	if err := services.CreateRoom(&room, userID.(uint)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, room)
}

func GetRoom(c *gin.Context) {
	roomID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid room ID")
		return
	}

	room, err := services.GetRoomByID(uint(roomID))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, room)
}

func DeleteRoom(c *gin.Context) {
	roomID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid room ID")
		return
	}

	if err := services.DeleteRoom(uint(roomID)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Room deleted successfully")
}

func JoinRoom(c *gin.Context) {
	roomID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid room ID")
		return
	}

	userID, _ := c.Get("userID")

	if err := services.JoinRoom(uint(roomID), userID.(uint)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "Joined room successfully")
}
