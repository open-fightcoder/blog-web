package models

import "testing"

func TestArticleTypeCreate(t *testing.T) {
	InitAllInTest()
	articleType:=&ArticleType{UserId:1,Index:1,ArticleTypeInfo:"123456"}
	if _, err := ArticleTypeCreate(articleType); err != nil {
		t.Error("Create() failed.Error:", err)
	}
}

func TestArticleTypeRemove(t *testing.T) {
	InitAllInTest()
	if err := ArticleTypeRemove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestArticleTypeUpdate(t *testing.T) {
	InitAllInTest()

	articleType:=&ArticleType{Id:2,UserId:1,Index:2,ArticleTypeInfo:"123456"}
	if err := ArticleTypeUpdate(articleType); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestArticleTypeGetById(t *testing.T) {
	InitAllInTest()
	_, err := ArticleTypeGetById(2)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}
}
