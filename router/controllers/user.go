package controllers

import (
	"net/http"

	"encoding/base64"

	error "github.com/blog-web/common/g"
	"github.com/blog-web/managers"
	"github.com/blog-web/models"
	"github.com/blog-web/router/controllers/base"
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.RouterGroup) {
	router.POST("user/login", httpHandlerLogin)
	router.POST("user/register", httpHandlerRegister)
	router.POST("user/changepasswd", httpHandlerChangePasswd)
	router.POST("user/changemess", httpHandlerChangeMess)
}

func httpHandlerLogin(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		panic(error.ParamError("参数格式错误,解析失败!"))
	}
	token, userId := managers.UserLogin(account.Email, account.Password)
	cookie := &http.Cookie{
		Name:     "token",
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, map[string]int64{
		"user_id": userId,
	})
}

func httpHandlerRegister(c *gin.Context) {
	type AccountRegister struct {
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
		UserName string `form:"user_name" json:"user_name"`
		NickName string `form:"nick_name" json:"nick_name"`
	}
	account := AccountRegister{}
	err := c.Bind(&account)
	if err != nil {
		panic(error.ParamError("参数格式错误,解析失败!"))
	}
	userId := managers.UserRegister(account.UserName, account.NickName, account.Email, account.Password)
	c.JSON(http.StatusOK, map[string]int64{
		"user_id": userId,
	})
}

func httpHandlerChangePasswd(c *gin.Context) {
	passwd := c.PostForm("passwd")
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	managers.UserChangePassWd(userId, passwd)
	c.JSON(http.StatusOK, "")
}

func httpHandlerChangeMess(c *gin.Context) {
	userId := base.UserId(c)
	if userId == 0 {
		panic(error.PrivError("您尚未登录!"))
	}
	param := models.User{}
	err := c.Bind(&param)
	if err != nil {
		panic(error.ParamError("参数格式错误,解析失败!"))
	}
	managers.UserChangeMess(userId, param.NickName, param.Sex, param.Blog, param.Git, param.Description, param.Birthday, param.DailyAddress, param.StatSchool, param.SchoolName)
	c.JSON(http.StatusOK, "")
}
