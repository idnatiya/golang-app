package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/models"
	"idnatiya.com/golang-app/utils"
)

type CreatePermissionType struct {
	Slug string `json:"slug" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func ListPermission(context *gin.Context) {
	var permissions []models.Permission
	if err := models.DB.Find(&permissions).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully get list permission",
		"data":    permissions,
	})
}

func CreatePermission(context *gin.Context) {
	var createPermissionType CreatePermissionType
	if err := context.ShouldBindJSON(&createPermissionType); err != nil {
		errors := utils.HandleValidationError(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Data is invalid",
			"errors":  errors,
		})
		return
	}

	permission := models.Permission{}
	permission.Slug = createPermissionType.Slug
	permission.Name = createPermissionType.Name

	if err := models.DB.Create(&permission).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Successfully create new permission",
		"data":    permission,
	})
}

type UpdatePermissionType struct {
	Slug string `json:"slug" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func UpdatePermission(context *gin.Context) {
	var updatePermissionType UpdatePermissionType
	if err := context.ShouldBindJSON(&updatePermissionType); err != nil {
		errors := utils.HandleValidationError(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Data is invalid",
			"errors":  errors,
		})
		return
	}

	var permission models.Permission
	if err := models.DB.First(&permission, context.Param("permissionID")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Permission not found",
		})
		return
	}

	permission.Slug = updatePermissionType.Slug
	permission.Name = updatePermissionType.Name

	if err := models.DB.Save(&permission).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully update permission",
		"data":    permission,
	})
}

func DeletePermission(context *gin.Context) {
	var permission models.Permission
	if err := models.DB.First(&permission, context.Param("permissionID")).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Permission not found",
		})
		return
	}

	models.DB.Delete(&permission)

	context.JSON(http.StatusNoContent, nil)
}
