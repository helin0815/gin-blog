package api

import (
	_ "github.com/gin-blog/faye-wong/model"
	"github.com/gin-gonic/gin"
)

type Apis struct {
	user *User
}

func NewAPIS(user *User) *Apis {
	return &Apis{
		user: user,
	}
}

func (as *Apis) SetupRoute(r *gin.Engine) {

	r.POST("/login", as.user.Login)
	r.Group("/api/v1")
	{
		r.GET("/hello", as.HelloWorld)
	}
}
