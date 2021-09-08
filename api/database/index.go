package database

import (
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
	r.GET("/fill-test-data", d.FillTestData)
	r.GET("/:dbname", d.List)
}
func (d *databaseManager) FillTestData(c *gin.Context) {
	t := c.Param("table")
	s := c.Query("size")
	tag := c.GetHeader("tag")
	c.JSON(200, gin.H{"table": t, "size": s, "tag": tag})
}
func (d *databaseManager) List(c *gin.Context) {
	t := c.Param("tablename")
	s := c.Query("size")
	c.JSON(200, gin.H{"table": t, "size": s})
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
