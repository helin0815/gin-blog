package service

import (
	"fmt"

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
func (u *User) AccountExist(name string) (bool, error) {
	var count int
	// 统计满足条件的记录数量
	err := u.db.Model(&model.Users{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count == 0 {
		// 没有数据
		return false, nil
	}
	// 存在数据
	return true, nil
}

func (u *User) AddAccount(user *model.Users) error {
	fmt.Println("user:", user.Name)
	return u.db.Create(user).Error
}
