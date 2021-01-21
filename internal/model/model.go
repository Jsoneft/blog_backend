package model

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/pkg/setting"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(Databasesettings *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local", Databasesettings.Username, Databasesettings.Password, Databasesettings.Host, Databasesettings.Port, Databasesettings.DBName, Databasesettings.Charset, Databasesettings.ParseTime)
	db, err := gorm.Open(Databasesettings.DBType, dsn)

	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(Databasesettings.MaxIdleConns)
	db.DB().SetMaxOpenConns(Databasesettings.MaxOpenConns)
	return db, nil
}
