package dao

import (
	"ginblog_backend/internal/model"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Model:         &model.Model{CreatedBy: param.CreatedBy},
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
	}
	return article.Create(d.engine)
}
