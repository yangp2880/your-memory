package db

// 初始化mysql
//func InitMysql(v *viper.Viper) *gorm.DB {
//	dsn := v.GetString("mysql.dsn")
//	var db *gorm.DB
//	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // 禁止自动复数
//		},
//	})
//	if err != nil {
//		log.Fatalf("无法连接到数据库: %v", err)
//	}
//	fmt.Println("mysql连接成功")
//	db = gormDb
//	if err != nil {
//		log.Fatal(err)
//		return nil
//	}
//	return db
//}
