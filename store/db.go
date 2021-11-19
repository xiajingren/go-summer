package store

import (
	"fmt"
	"log"

	"github.com/xiajingren/go-summer/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Conf.MySql.User,
		conf.Conf.MySql.Password,
		conf.Conf.MySql.Host,
		conf.Conf.MySql.Post,
		conf.Conf.MySql.Db_Name,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&User{})
}
