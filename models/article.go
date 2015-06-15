package models

import (
	"fmt"
	"log"
)

var (
	TBL_ARTICLES = "articles"
)

type articleModel struct{}

var _articleModel = articleModel{}

func ArticleModel() *articleModel {
	return &_articleModel
}

func (*articleModel) Add(title, content,kind,author,ctime string) error {
	sql := fmt.Sprintf("INSERT INTO "+TBL_ARTICLES+" SET title='%s',content='%s', kind='%s',author='%s',ctime='%s'", title, content,kind,author,ctime)
    log.Println("insert sql:",sql)
	_, err := DB.Query(sql)
	if nil != err {
		log.Println(err)
		return err
	}

	return nil
}
