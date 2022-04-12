package middleware

import "github.com/gin-gonic/gin"

func TokenMiddleware(c *gin.Context) {
	//token := c.GetHeader("token")
	//c.Cookie("token")
	c.Set("userId", "")
}
