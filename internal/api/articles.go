package api

import (
	"github.com/gin-blog/faye-wong/internal/service"
	"github.com/gin-gonic/gin"
)

type Articles struct {
	aSvc *service.Articles
}

func NewArticles(aSvc *service.Articles) *Articles {
	return &Articles{
		aSvc: aSvc,
	}
}

func (a Articles) List(c *gin.Context) {
	articlesList := a.aSvc.ArticlesTitleList()
	c.JSON(200, articlesList)
}
