package database

import (
	"github.com/clh021/crud-api/conf"
	"github.com/clh021/crud-api/mock"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DbConnection struct {
	tag  string
	conn *gorm.DB
}
type databaseManager []DbConnection

var database = []DbConnection{}

func New() *databaseManager {
	return &databaseManager{}
}
func (d *databaseManager) Register(r *gin.RouterGroup) {
	r.GET("/fill-test-data", d.fillTestData)
	r.GET("/:dbname", d.list)
}
func (d *databaseManager) fillTestData(c *gin.Context) {
	t := c.Param("table")
	s := c.Query("size")
	tag := c.GetHeader("tag")
	if len(tag) == 0 {
		tag = conf.GetFirstServer().Tag
	}
	db := d.getConnByTag(tag)
	m := mock.NewMock(db)
	m.ReSetAll()
	c.JSON(200, gin.H{"table": t, "size": s, "tag": tag})
}
func (d *databaseManager) list(c *gin.Context) {
	t := c.Param("tablename")
	s := c.Query("size")
	c.JSON(200, gin.H{"table": t, "size": s})
}
