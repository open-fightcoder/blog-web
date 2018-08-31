package managers

import (
	"strconv"

	error "github.com/blog-web/common/g"
	. "github.com/blog-web/common/store"
	"github.com/blog-web/models"
)

func UserFollow(followedId int64, followId string) {
	id, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	user, err := models.UserGetById(id)
	if err != nil {
		panic(error.DBError())
	}
	if user == nil {
		panic(error.ParamError("关注用户不存在"))
	}
	relation := &models.Relation{FollowId: id, FollowedId: followedId}
	_, err = OrmWeb.Insert(relation)
	if err != nil {
		panic(error.DBError("关注失败!"))
	}
}

func UserUnfollow(followedId int64, followId string) {
	id, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	relation, err := models.RelationGetByConds(id, followedId)
	if err != nil {
		panic(error.DBError())
	}
	if relation == nil {
		panic(error.ParamError("您尚未关注该用户!"))
	}
	_, err = OrmWeb.Id(relation.Id).Delete(&models.Relation{})
	if err != nil {
		panic(error.DBError("取消关注失败!"))
	}
}

func UserListidol(id int64) []*models.User {
	relationList := make([]*models.Relation, 0)

	err := OrmWeb.Cols("follow_id").Where("followed_id=?", id).Find(&relationList)
	if err != nil {
		panic(error.DBError())
	}
	userList := make([]*models.User, 0)
	for _, u := range relationList {
		user := new(models.User)
		has, err := OrmWeb.Id(u.FollowId).Get(user)
		if err != nil {
			panic(error.DBError())
		}
		if !has {
			userList = append(userList, nil)
		}
		userList = append(userList, user)
	}
	return userList
}

func UserListfans(id int64) []*models.User {
	relationList := make([]*models.Relation, 0)

	err := OrmWeb.Cols("followed_id").Where("follow_id=?", id).Find(&relationList)
	if err != nil {
		panic(error.DBError())
	}
	userList := make([]*models.User, 0)
	for _, u := range relationList {
		user := new(models.User)
		has, err := OrmWeb.Id(u.FollowedId).Get(user)
		if err != nil {
			panic(error.DBError())
		}
		if !has {
			userList = append(userList, nil)
		}
		userList = append(userList, user)
	}
	return userList
}
