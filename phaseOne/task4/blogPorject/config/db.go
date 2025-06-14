package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	viper.SetConfigName("config") // 配置文件名(不带扩展名)
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	user := viper.GetString("mysql.user")
	pass := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")
	charset := viper.GetString("mysql.charset")
	parseTime := viper.GetBool("mysql.parseTime")
	loc := viper.GetString("mysql.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		user, pass, host, port, dbname, charset, parseTime, loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有SQL
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	DB = db
	fmt.Println("数据库连接成功")
}
