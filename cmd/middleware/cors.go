package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		c.Writer.Header().Set("Connection", "keep-alive")

		c.Next()
	}
}
