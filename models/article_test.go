package models

import "testing"

func TestArticleCreate(t *testing.T) {
	InitAllInTest()
	article := &Article{Title:"123456",Content:"qwew",Status:1,ArticleType:1,Look:1,Favour:1,UserId:12}
	if _, err := ArticleCreate(article); err != nil {
		t.Error("Create() failed.Error:", err)
	}
}

func TestArticleRemove(t *testing.T) {
	InitAllInTest()
	if err := ArticleRemove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestArticleUpdate(t *testing.T) {
	InitAllInTest()

	article := &Article{Id:2,Title:"654321",Content:"qwew",Status:1,ArticleType:1,Look:1,Favour:1,UserId:12}
	if err := ArticleUpdate(article); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestArticleGetById(t *testing.T) {
	InitAllInTest()
	_, err := ArticleGetById(2)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}
}