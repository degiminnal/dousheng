package entity

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	var err error
	DB, err = gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		return err
	}
	// 创建数据表
	DB.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Comment{}, &Relation{})
	return nil
}
