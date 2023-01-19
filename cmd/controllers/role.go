package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/models"
	"idnatiya.com/golang-app/utils"
)

type CreateRoleRequest struct {
	Slug       string `json:"slug" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Permission []uint `json:"permissions"`
}

func ListRole(context *gin.Context) {
	var roles []models.Role
	if err := models.DB.Select("id", "slug").Preload("Permissions").Find(&roles).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully Create Role",
		"data":    roles,
	})
}

func CreateRole(context *gin.Context) {
	var createRoleRequest CreateRoleRequest
	if err := context.ShouldBindJSON(&createRoleRequest); err != nil {
		errors := utils.HandleValidationError(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Data is invalid",
			"errors":  errors,
		})
		return
	}

	var permissions []models.Permission
	for _, permissionID := range createRoleRequest.Permission {
		var permission models.Permission
		if err := models.DB.First(&permission, permissionID).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": "Data is invalid",
				"errors":  "Permission does`nt exists",
			})
			return
		}

		permissions = append(permissions, permission)
	}

	role := models.Role{
		Slug:        createRoleRequest.Slug,
		Name:        createRoleRequest.Name,
		Permissions: permissions,
	}

	if err := models.DB.Create(&role).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully Create Role Data",
		"data":    role,
	})
}
