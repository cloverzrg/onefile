package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(c *gin.Context) {
	//token := c.GetHeader("token")
	//c.Cookie("token")
	if c.FullPath() == "/login" {
		c.Next()
		return
	}
	if c.FullPath() == "/callback" {
		c.Next()
		return
	}
	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	userIdStr := userId.(string)
	if len(userIdStr) == 0 {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	c.Set("userId", userIdStr)
}
