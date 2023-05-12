package model

type Update struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}