package middleware

import (
	"github.com/cloverzrg/onefile/db"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func GetSessionsMiddleware() gin.HandlerFunc {
	store := gormsessions.NewStore(db.DB, true, []byte("onfile_123"))
	return sessions.Sessions("onefileSession", store)
}
