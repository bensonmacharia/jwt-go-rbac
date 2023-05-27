package controller

import (
	"bmacharia/jwt-go-rbac/model"
	"bmacharia/jwt-go-rbac/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// add Booking
func CreateBooking(c *gin.Context) {
	var input model.Booking
	var user_id = util.CurrentUser(c).ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.UserID != 0 {
		user_id = input.RoomID
	}
	booking := model.Booking{
		Status: "NOT PAID",
		UserID: user_id,
		RoomID: input.RoomID,
	}
	savedBooking, err := booking.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking done successfuly", "booking": savedBooking})
}

// get all Bookings
func GetBookings(c *gin.Context) {
	var Booking []model.Booking
	err := model.GetBookings(&Booking)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Booking)
}

// get logged in user bookings
func GetUserBookings(c *gin.Context) {
	var Booking model.Booking
	var user_id = util.CurrentUser(c).ID
	err := model.GetUserBookings(&Booking, user_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Booking)
}
