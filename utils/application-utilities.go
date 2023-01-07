package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserData(context *gin.Context) jwt.MapClaims {
	contextValue, _ := context.Get("userInfo")
	return contextValue.(jwt.MapClaims)
}
