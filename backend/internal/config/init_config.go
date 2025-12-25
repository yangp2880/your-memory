package config

import "github.com/spf13/viper"

func MysqlConfig() string {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	dsn := v.GetString("mysql.dsn")
	return dsn
}
