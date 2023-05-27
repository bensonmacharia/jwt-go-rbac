package model

import (
	"bmacharia/jwt-go-rbac/database"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Name     string `gorm:"not null;unique" json:"name"`
	Location string `gorm:"not null" json:"location"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// add a room
func (room *Room) Save() (*Room, error) {
	err := database.Db.Create(&room).Error
	if err != nil {
		return &Room{}, err
	}
	return room, nil
}

// get all rooms
func GetRooms(Room *[]Room) (err error) {
	err = database.Db.Find(Room).Error
	if err != nil {
		return err
	}
	return nil
}

// get room by id
func GetRoom(Room *Room, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Room).Error
	if err != nil {
		return err
	}
	return nil
}

// update room
func UpdateRoom(Room *Room) (err error) {
	database.Db.Save(Room)
	return nil
}
