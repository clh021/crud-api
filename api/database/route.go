package database

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	r.GET("/:dbname", List)
}
