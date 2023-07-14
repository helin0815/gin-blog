package api

import (
	"fmt"

	"github.com/gin-blog/faye-wong/internal/service"
	"github.com/gin-blog/faye-wong/model"
	"github.com/gin-gonic/gin"
)

type User struct {
	uSvc *service.User
}

func NewUser(uSvc *service.User) *User {
	return &User{uSvc: uSvc}
}
func (u *User) Login(c *gin.Context) {
	userM := new(model.Users)
	if err := c.ShouldBindJSON(userM); err != nil {
		fmt.Println("用户参数不对:", err.Error())
		return
	}
	fmt.Println("user name:", userM)
	isAdmin, err := u.uSvc.GetInfo(userM.Name, userM.Password)
	if err != nil {
		c.JSON(400, "登录失败")
	}
	if isAdmin {
		c.JSON(200, "登录成功")
	}

}
