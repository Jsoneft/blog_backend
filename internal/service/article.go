package service

import (
	"ginblog_backend/internal/dao"
	"ginblog_backend/internal/model"
	"ginblog_backend/pkg/app"
)

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagId         uint32 `form:"tag_id" binding:"required, gte=1"`
	Title         string `form:"title" binding:"required, min=2, max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binding:"required,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	CreatedBy     string `form:"created_by" binding:"required, max=100"`
	State         uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagId         uint32 `form:"tag_id" binding:"required, gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binging:"max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required, max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleByTagIdRequest struct {
	TagId uint32 `form:"tag_id" binding:"required, gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := s.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		CreatedBy:     param.CreatedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	err = s.dao.CreateArticleTag(article.ID, param.TagId, param.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := s.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	articleTag, err := s.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	tag, err := s.dao.GetTag(articleTag.TagId, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           &tag,
	}, nil
}

// GetArticleList return articles filtered by TagID.
func (s *Service) GetArticleList(param *ArticleListRequest, pager app.Pager) ([]*Article, int, error) {
	articleCount, err := s.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articleRows, err := s.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var articles []*Article
	for _, article := range articleRows {
		art := &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag: &model.Tag{
				Model: &model.Model{
					ID: article.TagID,
				},
				Name: article.TagName,
			},
		}
		articles = append(articles, art)
	}
	return articles, articleCount, nil
}

func (s Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := s.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		TagID:         param.TagId,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		ModifiedBy:    param.ModifiedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	err = s.dao.UpdateArticleTag(param.ID, param.TagId, param.ModifiedBy)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := s.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}
	err = s.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}
	return err
}
