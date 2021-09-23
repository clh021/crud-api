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
	// 创建测试表，填充测试数据
	// 未指定具体数据库则使用第一项配置
	r.GET("/fill-test-data", d.fillTestData)
	// 列举所有表
	r.GET("/list", d.list)
	// TODO: 分析表
	// TODO: 优化表
	// TODO: 检查表
	// TODO: 修复表
	// TODO: 清空表
	// TODO: 删除表
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
	tag := c.GetHeader("tag")
	if len(tag) == 0 {
		tag = conf.GetFirstServer().Tag
	}
	tables, err := d.getTableList(tag)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, gin.H{"table": tables})
	}
}
