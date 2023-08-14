package model

import "time"

type Users struct {
	UID      int    `json:"uid" gorm:"uid;type:int unsigned"`
	Name     string `json:"name" gorm:"name" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
	Mail     string `json:"mail" gorm:"mail" binding:"required"`
	URL      string `json:"url" gorm:"url"`
	// todo：如果下面不设置column:screenName的话，golang会把ScreenName解析成screen_name
	// 就会出现Unknown column 'screen_name' in 'field list' 的错误
	ScreenName string    `json:"screenName" gorm:"column:screenName" binding:"required"`
	Created    time.Time `json:"created" gorm:"created"`
	Activated  time.Time `json:"activated" gorm:"activated"`
	Logged     time.Time `json:"logged" gorm:"logged"`
	UserGroup  string    `json:"userGroup" gorm:"column:userGroup"`
	AuthCode   string    `json:"authCode" gorm:"column:authCode"`
}
type LoginUser struct {
	Name     string `json:"name" gorm:"name" binging:"required"`
	Password string `json:"password" gorm:"password" binging:"required"`
}
