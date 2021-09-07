package table

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TableManager struct {
	fields []string
}

func New() *TableManager {
	return &TableManager{}
}

func (t *TableManager) Register(r *gin.RouterGroup) {
	r.GET("/:tablename", t.List)
	r.POST("/:tablename", t.Create)
	r.GET("/:tablename/:primaryVal", t.Read)
	r.PUT("/:tablename/:primaryVal", t.Update)
	r.DELETE("/:tablename/:primaryVal", t.Delete)
}

func (t *TableManager) List(c *gin.Context) {
	tb := c.Param("tablename")
	s := c.Query("size")
	fmt.Printf("tablename: %s\n", tb)
	fmt.Printf("size: %s\n", s)
	respone := fmt.Sprintf("这里将显示 表 %s，size %s !", tb, s)
	// c.String(http.StatusOK, respone)
	c.JSON(200, gin.H{"message": respone, "table": t, "size": s})
}
func (t *TableManager) Create(c *gin.Context) {
	tb := c.Param("tablename")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}

func (t *TableManager) Delete(c *gin.Context) {
	tb := c.Param("tablename")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
func (t *TableManager) Read(c *gin.Context) {
	tb := c.Param("tablename")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
func (t *TableManager) Update(c *gin.Context) {
	tb := c.Param("tablename")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
