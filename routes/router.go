package routes

import (
	"myServer/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			log.InfoLog("get ping")
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return r
}
