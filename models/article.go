package models

import (
	. "github.com/blog-web/common/store"
)

type Article struct {
	Id int64 `form:"id" json:"id"`
	UserId int64 `form:"user_id" json:"user_id"`
	Title string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status int64 `form:"status" json:"status"`
	ArticleType int64 `form:"article_type" json:"article_type"`
	Look int64 `form:"look" json:"look"`
	Favour int64 `from:"favour" json:"favour"`
}

func ArticleCreate(article *Article) (int64, error) {
	return OrmWeb.Insert(article)
}

func ArticleRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Article{})
	return err
}

func ArticleUpdate(article *Article) error {
	_, err := OrmWeb.AllCols().ID(article.Id).Update(article)
	return err
}

func ArticleGetById(id int64) (*Article, error) {
	article := new(Article)
	has, err := OrmWeb.Id(id).Get(article)

	if err != nil{
		return nil, err
	}
	if !has{
		return nil,nil
	}
	return article, nil
}



