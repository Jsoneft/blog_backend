package dao

import (
	"ginblog_backend/internal/model"
	"ginblog_backend/pkg/app"
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

func (d Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
		State: state,
	}
	return article.Get(d.engine)
}

func (d Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	return article.Delete(d.engine)
}

func (d Dao) UpdateArticle(param *Article) error {
	article := model.Article{
		Model: &model.Model{ID: param.ID},
	}
	values := map[string]interface{}{
		"state":      param.State,
		"updated_by": param.ModifiedBy,
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)
}

func (d Dao) CountArticleListByTagID(tagID uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagId(d.engine, tagID)
}

func (d Dao) GetArticleListByTagID(tagID uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, tagID, pageOffset, pageSize)
}
