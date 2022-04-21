package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//链接数据库
func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=97301335 dbname=Shopsystem port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// 禁用复数形式
	//db.AutoMigrate(&models.Staff{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
