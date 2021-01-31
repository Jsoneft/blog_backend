package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (t ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	if err := db.Model(&t).Where("article_id = ? AND is_del = ?", t.ArticleId, 0).First(&articleTag).Error; err != nil && nil != gorm.ErrRecordNotFound {
		return ArticleTag{}, err
	}
	return articleTag, nil
}

func (t ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Model(&t).Where("tag_id = ? AND is_del = ?", t.TagId, 0).Find(&articleTags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (t ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Model(&t).Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(articleTags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (t ArticleTag) Create(db *gorm.DB) error {
	if err := db.Model(&t).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOne updates the row connected to the updated Article.
func (t ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&t).Where("article_id = ? AND is_del = ?", t.ID, 0).Limit(1).Updates(&values).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (t ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (t ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Model(&t).Where("article_id = ? AND is_del = ?", t.ArticleId, 0).Delete(&t).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
