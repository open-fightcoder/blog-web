package models

import (
	. "github.com/blog-web/common/store"
)

type Relation struct {
	Id         int64
	FollowId   int64 `form:"follow_id" json:"follow_id"`     //被关注者账户ID
	FollowedId int64 `form:"followed_id" json:"followed_id"` //关注者账户ID
}

func RelationGetByConds(followId int64, followedId int64) (*Relation, error) {
	relation := &Relation{}
	has, err := OrmWeb.Where("follow_id=? and followed_id=?", followId, followedId).Get(relation)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return relation, nil
}
