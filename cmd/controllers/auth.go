package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/models"
	"idnatiya.com/golang-app/cmd/services"
	"idnatiya.com/golang-app/cmd/types"
	"idnatiya.com/golang-app/utils"
)

func Register(c *gin.Context) {
	var registerType types.RegisterType
	if err := c.ShouldBindJSON(&registerType); err != nil {
		errors := utils.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Data is invalid",
			"errors":  errors,
		})
		return
	}

	authService := services.AuthServiceImpl{}

	user, err := authService.GetUserByEmail(&registerType.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "User is already registered",
		})
		return
	}

	user, err = authService.Register(&registerType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Succesfully Register User",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	var loginType types.LoginType
	if err := c.ShouldBindJSON(&loginType); err != nil {
		errors := utils.HandleValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Data is invalid",
			"errors":  errors,
		})
		return
	}

	authService := services.AuthServiceImpl{}

	token, err := authService.Attemp(&loginType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successfully",
		"token":   token,
	})
}

func RefreshToken(context *gin.Context) {
	userData := utils.GetUserData(context)

	var user models.User
	if err := models.DB.First(&user, userData["ID"]).Error; err != nil {
		context.JSON(http.StatusForbidden, gin.H{})
	}

	authService := services.AuthServiceImpl{}
	token, _ := authService.Login(&user)

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
