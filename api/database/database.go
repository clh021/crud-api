package database

import "github.com/jinzhu/gorm"

type DbConnection struct {
	tag  string
	conn *gorm.DB
}

type Database []DbConnection

func NewDatabase() Database {
	return []DbConnection{}
}

func NewConnection() {
	c := &DbConnection{
		tag:  "s",
		conn: gorm.Open("jaja"),
	}
	return c
}
func (dbc *DbConnection) List() {
	// gorm.Open("")
}
