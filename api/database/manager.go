package database

import (
	"log"
	"os"
	"time"

	"github.com/clh021/crud-api/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

// 将会在没有找到链接时自动链接
func (d *databaseManager) getConnByTag(tag string) *gorm.DB {
	for _, v := range database {
		if v.tag == tag {
			return v.conn
		}
	}
	return d.addTagConn(tag)
}

func (d *databaseManager) addTagConn(tag string) *gorm.DB {
	db, err := gorm.Open(d.getDialectorByTag(tag), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold: time.Second,   // 慢 SQL 阈值
				LogLevel:      logger.Silent, // 日志级别
				// IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				// Colorful:                  false,         // 禁用彩色打印
			},
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	database = append(database, DbConnection{
		tag:  tag,
		conn: db,
	})
	return db
}
