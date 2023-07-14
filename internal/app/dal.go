package app

import (
	"github.com/gin-blog/faye-wong/model"
	"github.com/jinzhu/gorm"
)

type Query struct {
	db *gorm.DB

	user model.Users
}
