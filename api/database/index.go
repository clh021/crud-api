package database

import (
	"fmt"
	"log"

	"github.com/clh021/crud-api/conf"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
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
	r.GET("/:dbname", d.List)
}
func (d *databaseManager) List(c *gin.Context) {
	t := c.Param("tablename")
	s := c.Query("size")
	fmt.Printf("tablename: %s\n", t)
	fmt.Printf("size: %s\n", s)
	respone := fmt.Sprintf("这里将显示 表 %s，size %s !", t, s)
	// c.String(http.StatusOK, respone)
	c.JSON(200, gin.H{"message": respone, "table": t, "size": s})
}
func (d *databaseManager) getDialectorByTag(tag string) gorm.Dialector {
	c := conf.GetServerByTag(tag)
	switch c.Type {
	case "mysql":
		return mysql.Open(c.Dsn)
	case "sqlite":
		return sqlite.Open(c.Dsn)
	case "sqlserver":
		return sqlserver.Open(c.Dsn)
	case "postgres":
		return postgres.Open(c.Dsn)
	}
	return mysql.Open(c.Dsn)
}
func (d *databaseManager) AddTagConn(tag string) {
	db, err := gorm.Open(d.getDialectorByTag(tag), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database = append(database, DbConnection{
		tag:  tag,
		conn: db,
	})
}