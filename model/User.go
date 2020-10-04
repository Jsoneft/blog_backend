package model

import (
	"ginblog_backend/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null " json:"username"`
	Password string `gorm:"type: varchar(20);not null " json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var data User
	db.Select("id").Where("username = ?", name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 添加用户
func CreatUser(data *User) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表 (带翻页)
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var cnt int64
	if username != "" {
		db.Select("id, username, role").Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where("username LIKE %", username+"%").Count(&cnt)
		return users, cnt
	}
	db.Select("id, username, role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&cnt)
	if err != nil {
		return users, 0
	}
	return users, cnt

}

//编辑用户

//删除用户
