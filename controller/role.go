package controller

import (
	"bmacharia/jwt-go-rbac/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// create Role
func CreateRole(c *gin.Context) {
	var Role model.Role
	c.BindJSON(&Role)
	err := model.CreateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get all Roles
func GetRoles(c *gin.Context) {
	var Role []model.Role
	err := model.GetRoles(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Role by id
func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Role model.Role
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// update Role
func UpdateRole(c *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Role)
	err = model.UpdateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}
