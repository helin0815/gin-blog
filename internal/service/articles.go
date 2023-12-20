package service

import "github.com/jinzhu/gorm"

type Articles struct {
	db *gorm.DB
}

func NewArticles(db *gorm.DB) *Articles {
	return &Articles{
		db: db,
	}
}

func (a Articles) ArticlesTitleList() []string {
	value, ok := a.db.Get("info")
	if !ok {
		return nil
	}
	if titleList, ok := value.([]string); ok {
		return titleList
	}

	return nil
}
