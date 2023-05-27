package model

import (
	"bmacharia/jwt-go-rbac/database"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	UserID uint   `gorm:"not null" json:"user_id"`
	RoomID uint   `gorm:"not null" json:"room_id"`
	Status string `gorm:"not null" json:"status"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Room   Room   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// add a booking
func (booking *Booking) Save() (*Booking, error) {
	err := database.Db.Create(&booking).Error
	if err != nil {
		return &Booking{}, err
	}
	return booking, nil
}

// get all bookings
func GetBookings(Booking *[]Booking) (err error) {
	err = database.Db.Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// get user bookings
func GetUserBookings(Booking *Booking, uid uint) (err error) {
	err = database.Db.Where("user_id = ?", uid).First(Booking).Error
	if err != nil {
		return err
	}
	return nil
}
