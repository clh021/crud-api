package opera

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	t := c.Param("tablename")
	s := c.Query("size")
	fmt.Printf("tablename: %s\n", t)
	fmt.Printf("size: %s\n", s)
	respone := fmt.Sprintf("这里将显示 表 %s，size %s !", t, s)
	// c.String(http.StatusOK, respone)
	c.JSON(200, gin.H{"message": respone, "table": t, "size": s})
}
