package api

import (
	"fmt"
	"path"
	"time"

	"github.com/gin-blog/faye-wong/internal/service"
	"github.com/gin-blog/faye-wong/model"
	utils "github.com/gin-blog/faye-wong/pkg"
	"github.com/gin-gonic/gin"
)

type User struct {
	uSvc *service.User
}

func NewUser(uSvc *service.User) *User {
	return &User{uSvc: uSvc}
}
func (u *User) Login(c *gin.Context) {
	userM := new(model.LoginUser)
	if err := c.ShouldBindJSON(userM); err != nil {
		fmt.Println("用户参数不对:", err.Error())
		utils.JSONServerError(c, utils.GenerateErrMsg("登录失败,用户参数异常：", err.Error()))
		return
	}
	fmt.Println("user name:", userM)
	isAdmin, err := u.uSvc.GetInfo(userM.Name, userM.Password)
	if err != nil {
		utils.JSONServerError(c, utils.GenerateErrMsg("登录失败：", err.Error()))
		return
	}
	if isAdmin {
		c.JSON(200, "登录成功")
		return
	}

}

func (u *User) Register(c *gin.Context) {
	userM := new(model.Users)
	if err := c.ShouldBindJSON(userM); err != nil {
		utils.JSONServerError(c, utils.GenerateErrMsg("注册失败,用户参数不足：", err.Error()))
		return
	}

	isExist, err := u.uSvc.AccountExist(userM.Name)
	if err != nil {
		utils.JSONServerError(c, utils.GenerateErrMsg("注册失败：", err.Error()))
		return
	}
	if isExist {
		utils.JSONOk(c, "注册失败，用户名已存在")
		return
	}
	u.setUserCommonInfo(userM)
	if err = u.uSvc.AddAccount(userM); err != nil {
		utils.JSONServerError(c, utils.GenerateErrMsg("注册失败2：", err.Error()))
		return
	}
	utils.JSONOk(c, "注册成功")
	return
}
func (u *User) setUserCommonInfo(userM *model.Users) {
	userM.Created = time.Now()
	userM.Activated = time.Now()
	userM.Logged = time.Now()
	userM.UserGroup = "visitor"
	userM.AuthCode = path.Join(userM.Name, "-", utils.GetCurrentTime())
}
