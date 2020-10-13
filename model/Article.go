package model

import (
	"ginblog_backend/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type int; not null" json:"cid"`
	Desc    string `gorm:"type varchar(200)" json:"desc"`
	Content string `gorm:"type longtext" json:"content"`
	Img     string `gorm:"type varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询单个文章
func GetArtInfo(id int )(Article, int){
	var art Article
	err := db.Model(&art).Where("id = ?", id).First(&art).Error
	// err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art,errmsg.ERROR_ART_NOT_EXIST
	}
	return art,errmsg.SUCCESS
}

// todo 查询分类下所有文章
func GetCateArt(id int,pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1)*pageSize).Where("id = ?", id).Find(&articles).Error
	if err != nil {
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}
	return articles, errmsg.SUCCESS
}

//  查询文章列表
func GetArticles(ArticleTitle string, pageSize int, pageNum int) ([]Article, int, int64) {
	var (
		articles []Article
		err      error
		cnt      int64
	)

	if ArticleTitle == "" {
		// 时间 从新到旧
		err = db.Order("Updated_At DESC").Preload("Category").Find(&articles).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error

		// Mysql: 'SELECT count(1) FROM `articles` WHERE `articles`.`deleted_at` IS NULL'
		// 统计数据库中 article 中有效的 条目， 并不是 所查询的 articles 的个数
		db.Model(&articles).Count(&cnt)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return articles, errmsg.SUCCESS, cnt
	}

	err = db.Order("Updated_At DESC").Preload("Category").Where("title LIKE ?", ArticleTitle+"%").Find(&articles).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&articles).Count(&cnt)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, cnt
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	err = db.Model(&article).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
