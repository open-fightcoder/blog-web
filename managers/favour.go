package managers

import (
	"strconv"

	error "github.com/blog-web/common/g"
	. "github.com/blog-web/common/store"
	"github.com/blog-web/models"
)

func ArticleLove(userId int64, articleIdStr string) {
	//TODO 检查id对应的文章是否存在
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	favour := &models.Favour{UserId: userId, ArticleId: articleId}
	_, err = OrmWeb.Insert(favour)
	if err != nil {
		panic(error.DBError("添加喜欢失败!"))
	}
}

func ArticleUnlove(userId int64, articleIdStr string) {
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	favour, err := models.FavourGetByConds(userId, int64(articleId))
	if err != nil {
		panic(error.DBError())
	}
	if favour == nil {
		panic(error.ParamError("请检查取消文章id!"))
	}
	_, err = OrmWeb.Id(favour.Id).Delete(&models.Favour{})
	if err != nil {
		panic(error.DBError("取消喜欢失败!"))
	}
}
