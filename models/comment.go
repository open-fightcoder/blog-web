package models

import (
	. "github.com/blog-web/common/store"
)

type Comment struct {
	Id        int64  `form:"id" json:"id"`             //Id
	ArticleId int64  `form:"blog_id" json:"blog_id"`   //blogId
	Content   string `form:"content" json:"content"`   //评论内容
	UserId    int64  `form:"user_id" json:"user_id"`   //评论人账户ID
	ReplyId   int64  `form:"reply_id" json:"reply_id"` //0 未被回复 >0 回复的commentId
}

func CommentGetById(id int64) (*Comment, error) {
	comment := new(Comment)
	has, err := OrmWeb.Id(id).Get(comment)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return comment, nil
}
