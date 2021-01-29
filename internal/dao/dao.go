package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(e *gorm.DB) *Dao {
	return &Dao{engine: e}
}
