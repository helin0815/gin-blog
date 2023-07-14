package model

import (
	"github.com/gin-blog/faye-wong/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", configs.MysqlDSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}
