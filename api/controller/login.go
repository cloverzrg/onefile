package controller

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.String(200, "here is onefile!")
}

func Login(c *gin.Context) {
	c.Redirect(302, "")
}
