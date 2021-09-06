package api

import (
	"fmt"

	"github.com/clh021/crud-api/api/conf"
	"github.com/clh021/crud-api/api/database"
	"github.com/clh021/crud-api/api/table"
	"github.com/clh021/crud-api/ui"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// ui
	r.Use(static.Serve("/", EmbedFolder(ui.Dist, "dist")))

	// table
	tb := r.Group("/table")
	table.Route(tb)

	// database
	db := r.Group("/db")
	database.Route(db)

	return r
}

func Main() {
	c := conf.Get()
	port := fmt.Sprintf(":%d", c.Port)
	r := GetRouter()
	r.Run(port)
	// http.ListenAndServe(port, http.FileServer(http.FS(web)))
}
