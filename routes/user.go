package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//register user

func RegisterUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	//hash password
	err = user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"err":     err.Error(),
		})

		c.Abort()
		return
	}

	//isertuser
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "bad request",
			"error":   insertUser.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
