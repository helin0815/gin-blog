package api

import (
	"net/http"

	_ "github.com/gin-blog/faye-wong/model"
	"github.com/gin-gonic/gin"
)

type Apis struct {
	user     *User
	articles *Articles
}

func NewAPIS(user *User, articles *Articles) *Apis {
	return &Apis{
		user:     user,
		articles: articles,
	}
}

// 对应的前端文件是：/opt/data01/testcode/vuepro

func (as *Apis) SetupRoute(r *gin.Engine) {
	r.Use(CorsMiddleware())
	r.GET("/", as.HelloWorld)
	r.POST("/login", as.user.Login)
	//r.Group("/api/v1")
	r.Group("/")
	{

		r.GET("/hello", as.HelloWorld)
		r.POST("/register", as.user.Register)
		r.GET("/articles", as.articles.List)
	}
}

// CorsMiddleware 解决跨域的方法
func CorsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
