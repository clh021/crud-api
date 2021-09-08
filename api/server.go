package api

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Register(*gin.RouterGroup)
}
type Server struct {
	engine *gin.Engine
}

func (m *Server) AddService(mod Service, routeGroupPath string) {
	r := m.engine.Group(routeGroupPath)
	mod.Register(r)
}
func (m *Server) AddMiddleware(middleware gin.HandlerFunc) {
	m.engine.Use(middleware)
}
func (m *Server) Run(addr string) {
	m.engine.Run(addr)
}
func (m *Server) Engine() *gin.Engine {
	return m.engine
}
func InitServer(ginMode string) *Server {
	if ginMode == gin.ReleaseMode {
		gin.DefaultWriter = ioutil.Discard
	}
	gin.SetMode(ginMode)
	return &Server{
		engine: gin.Default(),
	}
}
