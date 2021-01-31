package dao

import "ginblog_backend/internal/model"

func (d Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d Dao) GetArticleTagListByTID(tagId uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagId: tagId}
	return articleTag.ListByTID(d.engine)
}

func (d Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d Dao) CreateArticleTag(articleID, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createdBy},
		TagId:     tagId,
		ArticleId: articleID,
	}
	return articleTag.Create(d.engine)
}

func (d Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{UpdatedBy: modifiedBy},
		ArticleId: articleID,
		TagId:     tagID,
	}
	values := map[string]interface{}{
		"article_id": articleID,
		"tag_id":     tagID,
		"updated_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleId: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}
