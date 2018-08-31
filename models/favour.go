package models

import (
	. "github.com/blog-web/common/store"
)

type Favour struct {
	Id        int64
	UserId    int64 `form:"user_id" json:"user_id"`       //点赞用户Id
	ArticleId int64 `form:"article_id" json:"article_id"` //文章Id
}

func FavourGetByConds(userId int64, articleId int64) (*Favour, error) {
	favour := &Favour{}
	has, err := OrmWeb.Where("user_id=? and article_id=?", userId, articleId).Get(favour)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return favour, nil
}
