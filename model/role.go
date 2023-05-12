package model

import (
	"bmacharia/jwt-go-rbac/database"

	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

// Create a role
func CreateRole(Role *Role) (err error) {
	err = database.Db.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = database.Db.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *Role) (err error) {
	database.Db.Save(Role)
	return nil
}
