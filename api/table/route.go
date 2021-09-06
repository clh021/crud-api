package table

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	r.GET("/:tablename", List)
	r.POST("/:tablename", Create)
	r.GET("/:tablename/:primaryVal", Read)
	r.PUT("/:tablename/:primaryVal", Update)
	r.DELETE("/:tablename/:primaryVal", Delete)
}
