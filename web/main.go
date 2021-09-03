package web

import (
	"net/http"

	"github.com/clh021/crud-api/web/opera"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome for you!")
	})
	r.GET("/:tablename", opera.List)
	r.POST("/:tablename", opera.Create)
	r.GET("/:tablename/:primaryVal", opera.Read)
	r.PUT("/:tablename/:primaryVal", opera.Update)
	r.DELETE("/:tablename/:primaryVal", opera.Delete)
	return r
}
func Main() {
	loadConfig()
	r := GetRouter()
	// TODO port set by cobra default
	r.Run(":8000")
}
