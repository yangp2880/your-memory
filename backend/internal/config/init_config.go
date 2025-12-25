package config

import "github.com/spf13/viper"

// 初始化配置文件
func InitConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return v
}

//func MysqlConfig(v *viper.Viper) string {
//	dsn := v.GetString("mysql.dsn")
//	return dsn
//}
//
//func PostgresConfig(v *viper.Viper) string {
//	dsn := v.GetString("postgres.dsn")
//	return dsn
//}
