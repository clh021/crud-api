package api

import (
	"fmt"
	"net/http"

	"github.com/clh021/crud-api/api/opera"

	"github.com/gin-gonic/gin"
)

func favicon(c *gin.Context) {
	c.Header("Content-Type", "image/x-icon")
	c.Header("Cache-Control", "public, max-age=7776000")
	c.String(http.StatusOK, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}
func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome for you!")
	})
	r.GET("/favicon.ico", favicon) // 包装进入静态资源，加入 favicon.ico
	// 后期会考虑加入请求生成器，列举出所有数据库，表格，通过勾选操作查询
	// 加入 表格优化操作，清空操作
	r.GET("/:tablename", opera.List)
	r.POST("/:tablename", opera.Create)
	r.GET("/:tablename/:primaryVal", opera.Read)
	r.PUT("/:tablename/:primaryVal", opera.Update)
	r.DELETE("/:tablename/:primaryVal", opera.Delete)
	return r
}

var Conf = loadConfig()

func Main() {
	r := GetRouter()
	r.Run(fmt.Sprintf(":%d", Conf.Port))
}
