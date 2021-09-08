package api

import (
	"fmt"

	"github.com/clh021/crud-api/api/database"
	"github.com/clh021/crud-api/api/table"
	"github.com/clh021/crud-api/conf"
	"github.com/clh021/crud-api/ui"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func EngineServer(ginMode string) (*Server, int32) {
	c := conf.Get()
	s := InitServer(ginMode)
	s.AddService(database.New(), "/db")
	s.AddService(table.New(), "/table")
	s.AddMiddleware(static.Serve("/", EmbedFolder(ui.Dist, "dist")))
	return s, c.Port
}

func Main() {
	s, p := EngineServer(gin.ReleaseMode)
	port := fmt.Sprintf(":%d", p)
	s.Run(port)
	// http.ListenAndServe(port, http.FileServer(http.FS(web)))
}
