package managers

import (
	"github.com/blog-web/common/components"
	error "github.com/blog-web/common/g"
	"github.com/blog-web/models"
)

func UserLogin(email string, password string) (string, int64) {
	account, err := models.UserGetByEmail(email)
	if err != nil {
		panic(error.DBError())
	}
	if account == nil {
		panic(error.ConflictError("邮箱不存在"))
	}
	passwd := account.Password
	if passwd != components.MD5Encode(password) {
		panic(error.ConflictError("密码错误"))
	}
	user, err := models.UserGetByAccountId(account.Id)
	if err != nil {
		panic(error.DBError())
	}
	token, err := components.CreateToken(user.Id)
	if err != nil {
		panic(error.ServerError())
	}
	return token, user.Id
}

func UserRegister(userName string, nickName string, email string, password string) int64 {
	//加锁
	hasAccount, err := models.UserGetByEmail(email)
	if err != nil {
		panic(error.DBError())
	}
	if hasAccount != nil {
		panic(error.ConflictError("邮箱已存在"))
	}
	hasUser, err := models.UserGetByUserName(userName)
	if err != nil {
		panic(error.DBError())
	}
	if hasUser != nil {
		panic(error.ConflictError("用户名已存在"))
	}
	account := &models.Account{Email: email, Password: password}
	user := &models.User{UserName: userName, NickName: nickName, Avator: "http://xupt1.fightcoder.com:9001/image/default.png"}
	userId, err := models.UserAdd(account, user)
	if err != nil {
		panic(error.DBError("创建用户失败"))
	}
	return userId
}
