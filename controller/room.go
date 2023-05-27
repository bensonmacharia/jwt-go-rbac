package controller

import (
	"bmacharia/jwt-go-rbac/model"
	"bmacharia/jwt-go-rbac/util"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// add Room
func CreateRoom(c *gin.Context) {
	var input model.Room

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room := model.Room{
		Name:     input.Name,
		Location: input.Location,
		UserID:   util.CurrentUser(c).ID,
	}
	savedRoom, err := room.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room added successfuly", "room": savedRoom})
}

// get Rooms
func GetRooms(c *gin.Context) {
	var Room []model.Room
	err := model.GetRooms(&Room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Room)
}

// get Room by id
func GetRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Room model.Room
	err := model.GetRoom(&Room, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Room)
}

// update Room
func UpdateRoom(c *gin.Context) {
	var Room model.Room
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetRoom(&Room, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Room)
	err = model.UpdateRoom(&Room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfuly", "room": Room})
}
