package service

import (
	"github.com/gin-blog/faye-wong/model"
	"github.com/jinzhu/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) GetInfo(name, passwd string) (bool, error) {
	var auth model.Users
	db := u.db.Where("name = ?", name)
	err := db.First(&auth).Error
	if err != nil {
		return false, err
	}
	if auth.Password == passwd {
		return true, nil
	}
	return false, nil

}
