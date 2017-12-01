package models

import (
	"github.com/zanjs/y-mugg-v2/db"
)

//CreateTable user
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{}, &Article{}, &Product{}, &Wareroom{}, &Record{}, &Inventory{})
	return nil
}
