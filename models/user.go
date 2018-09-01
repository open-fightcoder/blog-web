package models

import (
	"errors"

	"github.com/blog-web/common/components"
	. "github.com/blog-web/common/store"
)

type User struct {
	Id           int64  `form:"id" json:"id"`
	AccountId    int64  `form:"account_id" json:"account_id"`       //账号Id
	UserName     string `form:"user_name" json:"user_name"`         //用户名
	NickName     string `form:"nick_name" json:"nick_name"`         //昵称
	Sex          string `form:"sex" json:"sex"`                     //性别
	Avator       string `form:"avator" json:"avator"`               //头像
	Blog         string `form:"blog" json:"blog"`                   //博客地址
	Git          string `form:"git" json:"git"`                     //Git地址
	Description  string `form:"description" json:"description"`     //个人描述
	Birthday     string `form:"birthday" json:"birthday"`           //生日
	DailyAddress string `form:"daily_address" json:"daily_address"` //日常所在地：省、市
	StatSchool   string `form:"stat_school" json:"stat_school"`     //当前就学状态(小学及以下、中学学生、大学学生、非在校生)
	SchoolName   string `form:"school_name" json:"school_name"`     //学校名称
}

func UserGetById(id int64) (*User, error) {
	user := new(User)
	has, err := OrmWeb.Id(id).Get(user)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func UserGetByUserName(userName string) (*User, error) {
	user := new(User)
	has, err := OrmWeb.Where("user_name = ?", userName).Get(user)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func UserGetByAccountId(accountId int64) (*User, error) {
	user := new(User)
	has, err := OrmWeb.Where("account_id = ?", accountId).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func UserUpdate(user *User) error {
	_, err := OrmWeb.AllCols().ID(user.Id).Update(user)
	return err
}

type Account struct {
	Id        int64
	Email     string `form:"email" json:"email"`       //邮箱
	Password  string `form:"password" json:"password"` //密码
	Phone     string //手机号
	QqId      string //用于QQ第三方登录
	GithubId  string //Github第三方登录
	WeichatId string //weichat第三方登录
}

func AccountUpdate(account *Account) error {
	_, err := OrmWeb.AllCols().ID(account.Id).Update(account)
	return err
}

func GetAccountByUserId(userId int64) (*Account, error) {
	user := new(User)
	has, err := OrmWeb.Id(userId).Get(user)
	if err != nil || !has {
		return nil, errors.New("DB Error")
	}
	account := new(Account)
	has, err = OrmWeb.Id(user.AccountId).Get(account)
	if err != nil || !has {
		return nil, errors.New("DB Error")
	}
	return account, nil
}

func UserAdd(account *Account, user *User) (int64, error) {
	session := OrmWeb.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return 0, err
	}
	account.Password = components.MD5Encode(account.Password)
	_, err = session.Insert(account)
	if err != nil {
		session.Rollback()
		return 0, err
	}
	user.AccountId = account.Id
	_, err = session.Insert(user)
	if err != nil {
		session.Rollback()
		return 0, err
	}
	err = session.Commit()
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func UserGetByEmail(email string) (*Account, error) {
	account := new(Account)
	has, err := OrmWeb.Where("email=?", email).Get(account)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return account, nil
}
