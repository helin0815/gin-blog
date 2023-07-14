package model

import "time"

type Users struct {
	UID        int       `json:"uid" gorm:"uid;type:int unsigned"`
	Name       string    `json:"name" gorm:"name"`
	Password   string    `json:"password" gorm:"password"`
	Mail       string    `json:"mail" gorm:"mail"`
	URL        string    `json:"url" gorm:"url"`
	ScreenName string    `json:"screenName" gorm:"screenName"`
	Created    time.Time `json:"created" gorm:"created"`
	Activated  time.Time `json:"activated" gorm:"activated"`
	Logged     time.Time `json:"logged" gorm:"logged"`
	UserGroup  string    `json:"userGroup" gorm:"userGroup"`
	AuthCode   string    `json:"authCode" gorm:"authCode"`
}
