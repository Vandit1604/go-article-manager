package models

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	// check if length of both the articlelist is same
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// checking if each value is same
	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
