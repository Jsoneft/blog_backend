package model

import "gorm.io/gorm"

type ArticleTag struct {
	gorm.Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
