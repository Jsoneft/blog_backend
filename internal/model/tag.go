package model

import (
	"ginblog_backend/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (a Tag) Count(db *gorm.DB) (int, error) {
	var res int
	if a.Name != "" {
		db = db.Where("Name = ?", a.Name)
	}
	db = db.Where("State = ?", a.State)
	if err := db.Model(&Tag{}).Where("is_del = ?", 0).Count(&res).Error; err != nil {
		return 0, err
	}
	return res, nil
}

func (a Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var res []*Tag
	if pageOffset >= 0 && pageSize >= 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Name != "" {
		db = db.Where("Name = ?", a.Name)
	}
	db = db.Where("State = ?", a.State)
	if err := db.Where("is_del = ?", 0).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (a Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", a.ID, 0, a.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (a Tag) Create(db *gorm.DB) error {
	return db.Model(&Tag{}).Create(&a).Error
}

func (a Tag) Delete(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}

func (a Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
