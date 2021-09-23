package database

import (
	"database/sql"
	"fmt"
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

func (d *databaseManager) getDialectorByTag(tag string) (gorm.Dialector, error) {
	c := conf.GetServerByTag(tag)
	switch c.Type {
	case "mysql":
		return mysql.Open(c.Dsn), nil
	case "sqlite":
		return sqlite.Open(c.Dsn), nil
	case "sqlserver":
		return sqlserver.Open(c.Dsn), nil
	case "postgres":
		return postgres.Open(c.Dsn), nil
	}
	return nil, fmt.Errorf("not support Database for %s on %s", c.Type, c.Tag)
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
	log.Println("addTagConn:", tag)
	di, err := d.getDialectorByTag(tag)
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(di, &gorm.Config{
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

func (d *databaseManager) getTableList(tag string) ([]Table, error) {
	c := conf.GetServerByTag(tag)
	var r *sql.Rows
	var t []Table
	err := fmt.Errorf("not support Database for %s on %s", c.Type, c.Tag)
	switch c.Type {
	case "mysql":
		db := d.getConnByTag(tag)
		r, err = db.Debug().Table("information_schema.tables").Not(map[string]interface{}{"table_schema": []string{"performance_schema", "information_schema", "mysql"}}).Rows()
		if err != nil {
			log.Fatalln(err)
		}
		var cs []string
		cs, err = r.Columns()
		defer r.Close()
		log.Println("cs:", cs)
		var st MysqlTable
		for r.Next() {
			log.Println("r:", r)
			db.ScanRows(r, &st)
		}
		log.Println("t:", st)
	case "sqlite":
		db := d.getConnByTag(tag)
		r, err = db.Debug().Table("sqlite_master").Where("type = ?", "table").Rows()
		if err != nil {
			log.Fatalln(err)
		}
		var cs []string
		cs, err = r.Columns()
		defer r.Close()
		log.Println("cs:", cs)
		var st SqliteTable
		for r.Next() {
			// ScanRows 方法用于将一行记录扫描至结构体
			log.Println("r:", r)
			db.ScanRows(r, &st)
		}
		log.Println("t:", st)
	case "sqlserver":
	case "postgres":
		break
	}
	return t, err
}
