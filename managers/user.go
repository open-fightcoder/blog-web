package managers

import (
	"github.com/blog-web/common/components"
	error "github.com/blog-web/common/g"
	"github.com/blog-web/models"
)

func UserChangeMess(userId int64, NickName string, Sex string, Blog string, Git string, Description string, Birthday string, DailyAddress string, StatSchool string, SchoolName string) {
	userMess, err := models.UserGetById(userId)
	if err != nil {
		panic(error.DBError())
	}
	userMess.NickName = NickName
	userMess.Sex = Sex
	userMess.Blog = Blog
	userMess.Git = Git
	userMess.Description = Description
	userMess.Birthday = Birthday
	userMess.DailyAddress = DailyAddress
	userMess.StatSchool = StatSchool
	userMess.SchoolName = SchoolName
	err = models.UserUpdate(userMess)
	if err != nil {
		panic(error.DBError("修改个人信息失败!"))
	}
}

func UserChangePassWd(userId int64, password string) {
	account, err := models.GetAccountByUserId(userId)
	if err != nil {
		panic(error.DBError())
	}
	account.Password = components.MD5Encode(password)
	err = models.AccountUpdate(account)
	if err != nil {
		panic(error.DBError("修改密码失败!"))
	}
}

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
