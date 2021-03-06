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
	// 指定表的数据分页
	r.GET("/:name", t.List)
	// 指定表，创建记录
	r.POST("/:name", t.Create)
	// 指定表，读取记录
	r.GET("/:name/:pkVal", t.Read)
	// 指定表，修改记录
	r.PUT("/:name/:pkVal", t.Update)
	// 指定表，删除记录
	r.DELETE("/:name/:pkVal", t.Delete)
}

func (t *TableManager) List(c *gin.Context) {
	tb := c.Param("name")
	s := c.Query("size")
	fmt.Printf("name: %s\n", tb)
	fmt.Printf("size: %s\n", s)
	respone := fmt.Sprintf("这里将显示 表 %s，size %s !", tb, s)
	// c.String(http.StatusOK, respone)
	c.JSON(200, gin.H{"message": respone, "table": t, "size": s})
}
func (t *TableManager) Create(c *gin.Context) {
	tb := c.Param("name")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}

func (t *TableManager) Delete(c *gin.Context) {
	tb := c.Param("name")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
func (t *TableManager) Read(c *gin.Context) {
	tb := c.Param("name")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
func (t *TableManager) Update(c *gin.Context) {
	tb := c.Param("name")
	respone := fmt.Sprintf("这里将显示 表 %s !", tb)
	c.JSON(200, gin.H{"message": respone, "table": tb})
}
