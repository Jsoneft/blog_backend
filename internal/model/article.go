package model

import (
	"ginblog_backend/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Model(&Article{}).Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
