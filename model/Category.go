package model

import (
	"ginblog_backend/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
}

// 查询分类名是否存在
func CheckCategoryName(CategoryName string) int {
	var data Category
	db.Select("id").Where("name = ?", CategoryName).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_EXIST
	}
	return errmsg.SUCCESS
}

// 添加分类
func CreateCategory(data *Category) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategories(CateName string, pageSize int, pageNum int) ([]Category, int64) {
	var categories []Category
	var cnt int64
	if CateName != "" {
		db.Select("id,name").Where("name LIKE ?", CateName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories)
		db.Model(&categories).Where("name LIKE ?", CateName+"%").Count(&cnt)
		return categories, cnt
	}
	db.Select("id, name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories)
	db.Model(&categories).Count(&cnt)
	return categories, cnt
}

// 编辑分类
func EditCategory(id int, data *Category)int  {
	var category Category
	err = db.Model(&category).Where("id = ?", id).Updates(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 删除分类

func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?",id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}


// todo 查询分类下的所有文章