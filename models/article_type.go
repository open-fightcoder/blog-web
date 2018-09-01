package models

import (
	. "github.com/blog-web/common/store"
)

type ArticleType struct {
	Id int64 `form:"id" json:"id"`
	UserId int64 `form:"user_id" json:"user_id"`
	Index int64 `form:"index" json:"index"`
	ArticleTypeInfo string `form:"article_type_info" json:"article_type_info"`
}

func ArticleTypeCreate(articleType *ArticleType) (int64, error) {
	return OrmWeb.Insert(articleType)
}

func ArticleTypeRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&ArticleType{})
	return err
}

func ArticleTypeUpdate(articleType *ArticleType) error {
	_, err := OrmWeb.AllCols().ID(articleType.Id).Update(articleType)
	return err
}

func ArticleTypeGetById(id int64) (*ArticleType, error) {
	articleType := new(ArticleType)
	has, err := OrmWeb.Id(id).Get(articleType)

	if err != nil {
		return nil, err
	}
	if !has{
		return nil,nil
	}
	return articleType, nil
}
