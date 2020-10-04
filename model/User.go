package model

import (
	"encoding/base64"
	"ginblog_backend/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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

//编辑用户 (不能修改密码)
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//加密
func ScryptPw(password string) (string, error) {
	const KeyLen = 10
	salt := []byte{114, 5, 71, 22, 41, 47, 81, 222}
	HashPw, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw, nil
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.Password, err = ScryptPw(user.Password)
	return err
}
