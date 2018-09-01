package managers

import (
	"strconv"

	error "github.com/blog-web/common/g"
	. "github.com/blog-web/common/store"
	"github.com/blog-web/models"
)

func CommentAdd(userId int64, blogIdStr string, content string) {
	blogId, err := strconv.ParseInt(blogIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	//TODO 添加 ArticleGetById 函数
	//blog, err := models.ArticleGetById(blogId)
	//if err != nil {
	//	panic(error.DBError())
	//}
	//if blog == nil {
	//	panic(error.ParamError("评论的文章不存在"))
	//}
	comment := &models.Comment{ArticleId: blogId, Content: content, UserId: userId, ReplyId: 0}
	_, err = OrmWeb.Insert(comment)
	if err != nil {
		panic(error.DBError("评论失败!"))
	}
}

func CommentReply(userId int64, commentIdStr string, blogIdStr string, content string) {
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	blogId, err := strconv.ParseInt(blogIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	//TODO 添加 ArticleGetById 函数
	//blog, err := models.ArticleGetById(blogId)
	//if err != nil {
	//	panic(error.DBError())
	//}
	//if blog == nil {
	//	panic(error.ParamError("评论的文章不存在"))
	//}
	com, err := models.CommentGetById(commentId)
	if err != nil {
		panic(error.DBError())
	}
	if com == nil {
		panic(error.ParamError("回复的评论不存在!"))
	}
	if com.ReplyId != 0 {
		panic(error.ParamError("回复的评论不存在!"))
	}
	comment := &models.Comment{ArticleId: blogId, Content: content, UserId: userId, ReplyId: commentId}
	_, err = OrmWeb.Insert(comment)
	if err != nil {
		panic(error.DBError("评论失败!"))
	}
}

func CommentDelete(userId int64, commentIdStr string) {
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	com, err := models.CommentGetById(commentId)
	if err != nil {
		panic(error.DBError())
	}
	if com == nil {
		panic(error.ParamError("评论不存在!"))
	}
	if com.UserId != userId {
		panic(error.PrivError("您无权操作!"))
	}
	if com.ReplyId == 0 {
		_, err = OrmWeb.Where("reply_id = ?", com.Id).Delete(&models.Comment{})
		if err != nil {
			panic(error.DBError("删除失败!"))
		}
	}
	_, err = OrmWeb.Id(com.Id).Delete(&models.Comment{})
	if err != nil {
		panic(error.DBError("删除失败!"))
	}
}
