package model

import (
	"fmt"
	"ginblog_backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.Dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库出现错误，请检查对应参数", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取DBcopy 出错", err)
	}
	err = db.AutoMigrate(&User{}, &Article{}, &Category{})
	if err != nil{
		fmt.Println("数据库迁移出错",err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//err = sqlDB.Close()
	//if err != nil{
	//	fmt.Println("数据库关闭出错",err)
	//}

}
