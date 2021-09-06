package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/clh021/crud-api/api/opera"
	"github.com/clh021/crud-api/ui"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}
func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
func GetRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	// ui
	r.Use(static.Serve("/", EmbedFolder(ui.Dist, "dist")))

	// api
	api := r.Group("/api")
	{
		api.GET("/:tablename", opera.List)
		api.POST("/:tablename", opera.Create)
		api.GET("/:tablename/:primaryVal", opera.Read)
		api.PUT("/:tablename/:primaryVal", opera.Update)
		api.DELETE("/:tablename/:primaryVal", opera.Delete)
	}
	return r
}

var Conf = loadConfig()

func Main() {
	port := fmt.Sprintf(":%d", Conf.Port)
	r := GetRouter()
	r.Run(port)
	// http.ListenAndServe(port, http.FileServer(http.FS(web)))
}
