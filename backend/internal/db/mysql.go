package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysql(dsn string) *gorm.DB {
	var db *gorm.DB
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止自动复数
		},
	})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	log.Printf("连接成功")
	db = gormDb
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
