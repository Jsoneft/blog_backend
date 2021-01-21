package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}
