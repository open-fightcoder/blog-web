package models

import "testing"

func TestArticleCreate(t *testing.T) {
	InitAllInTest()
	article := &Article{Title:"123456",UserId:12,Content:"qwew",Status:1,Type:1,Look:1,favour:1}
	if _, err := ArticleCreate(article); err != nil {
		t.Error("Create() failed.Error:", err)
	}
}