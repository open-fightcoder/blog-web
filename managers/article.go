package managers

import (
	"github.com/blog-web/models"
	error "github.com/blog-web/common/g"
	"strconv"
)

func ArticleTypeWriter(userId int64,indexStr,articleTypeInfo string){

	index, err := strconv.ParseInt(indexStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}

	articleType:=models.ArticleType{UserId:userId,Index:index,ArticleTypeInfo:articleTypeInfo}

	_, err = models.ArticleTypeCreate(&articleType)
	if err != nil {
		panic(error.DBError("添加文章类型失败!"))
	}
}
func ArticleTypeReader(articleTypeIdStr string) *models.ArticleType{
	articleTypeId, err := strconv.ParseInt(articleTypeIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	articleType, err := models.ArticleTypeGetById(articleTypeId)
	if err != nil {
		panic(error.DBError("获取文章类型失败!"))
	}
	return articleType
}
func ArticleWriter(userId int64,title,content,statusStr,articleTypeIdStr,lookStr,favourStr string){
	articleTypeId, err := strconv.ParseInt(articleTypeIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	look, err := strconv.ParseInt(lookStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	status, err := strconv.ParseInt(statusStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	favour, err := strconv.ParseInt(favourStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}

	article:=models.Article{UserId:userId,Title:title,Content:content,Status: status,ArticleType:articleTypeId,Look:look,Favour:favour}

	_, err = models.ArticleCreate(&article)
	if err != nil {
		panic(error.DBError("添加文章失败!"))
	}
}

func ArticleReader(articleTypeIdStr string) *models.Article{
	articleTypeId, err := strconv.ParseInt(articleTypeIdStr, 10, 64)
	if err != nil {
		panic(error.ParamError("参数必须为数字!"))
	}
	article, err := models.ArticleGetById(articleTypeId)
	if err != nil {
		panic(error.DBError("获取文章失败!"))
	}

	return article

}
