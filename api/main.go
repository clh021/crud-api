package api

import (
	"fmt"

	"github.com/clh021/crud-api/api/database"
	"github.com/clh021/crud-api/conf"
	"github.com/clh021/crud-api/ui"
	"github.com/gin-gonic/contrib/static"
)

func Main() {
	c := conf.Get()
	port := fmt.Sprintf(":%d", c.Port)
	s := InitServer()
	s.AddService(database.New(), "/db")
	// s.AddService(table.New(), "/table")
	s.AddMiddleware(static.Serve("/", EmbedFolder(ui.Dist, "dist")))
	s.Run(port)
	// http.ListenAndServe(port, http.FileServer(http.FS(web)))
}
