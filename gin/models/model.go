package models

import (
	"fmt"
	"gin/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	CreatedOn  int `json:"created_on" column:"created_on"`
	ModifiedOn int `json:"modified_on" column:"modified_on"`
	DeletedOn  int `json:"deleted_on"  column:"deleted_on"`
}

func Setup() {
	var err error
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("connect db server failed: %v", err)
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(time.Second * 5)
}
