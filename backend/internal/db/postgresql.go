package db

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 示例模型（普通表）
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"unique;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 如果你想用 pgvector 的 vector 类型（推荐安装 pgvector-go）
/*
import "github.com/pgvector/pgvector-go"

type Document struct {
    ID       uint
    Content  string
    Embedding pgvector.Vector `gorm:"type:vector(1536)"` // 例如 OpenAI 1536 维向量
}
*/

func InitPostgresql(v *viper.Viper) *gorm.DB {
	// PostgreSQL 连接字符串（DSN）
	// 根据你的 docker-compose 配置修改
	dsn := v.GetString("psql.dsn")

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败: ", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Ping 数据库失败: ", err)
	}
	fmt.Println("成功连接到 PostgreSQL！")

	// 自动迁移（创建表结构）
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("自动迁移失败: ", err)
	}
	fmt.Println("表迁移完成")

	// 关闭连接（程序结束时）
	sqlDB.Close()
	return db
}
