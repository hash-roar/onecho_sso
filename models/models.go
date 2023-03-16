package models

import (
	"fmt"
	"log"
	"onecho_sso_backend/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s  password=%s sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.DbName,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Password)
	db, err = gorm.Open(setting.DatabaseSetting.Type, dsn)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})

}
