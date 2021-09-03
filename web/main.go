package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome for you!")
	})
	return r
}
func Main() {
	r := GetRouter()
	// TODO port set by cobra default
	r.Run(":8000")
}
