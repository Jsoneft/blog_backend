package model

import (
	"fmt"
	"ginblog_backend/global"
	"ginblog_backend/pkg/setting"
	"time"

	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        uint32     `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"  sql:"index"`
	IsDel     uint8      `json:"is_del"`
}

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
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
	// 注册回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallBack)
	db.DB().SetMaxIdleConns(Databasesettings.MaxIdleConns)
	db.DB().SetMaxOpenConns(Databasesettings.MaxOpenConns)
	db.AutoMigrate(&Tag{}, &Article{}, &ArticleTag{})
	return db, nil
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updateTimeField.IsBlank {
				_ = updateTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when creating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", time.Now())
		}
	}
}

func deleteCallBack(scope *gorm.Scope) {
	var extraOption string
	if !scope.HasError() {
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
	}
	deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedOn")
	isDelField, hasIsDelField := scope.FieldByName("IsDel")
	if !scope.Search.Unscoped && hasDeletedAtField && hasIsDelField {
		// 如果 有 `Is_deleted` 和 `Deleted_at` 这两列 并且该scope 还有效的时候会执行软删除
		now := time.Now()
		scope.Raw(fmt.Sprintf("UPDATE %v SET %v=%v, %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deletedAtField.DBName),
			scope.AddToVars(now),
			scope.Quote(isDelField.DBName),
			scope.AddToVars(1),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	} else {
		// 硬删除
		scope.Raw(fmt.Sprintf("DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	}

}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return str
}
