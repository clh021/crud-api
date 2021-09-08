package api

import "github.com/gin-gonic/gin"

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
func InitServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}
