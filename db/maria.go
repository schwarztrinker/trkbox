package db

import (
	"log"
	"strconv"

	"github.com/schwarztrinker/trkbox/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var UsersDB Users
var maria *gorm.DB

func InitMariaDB() {
	dsn := conf.Conf.DB_USER + ":" + conf.Conf.DB_PASSWORD + "@tcp(" + conf.Conf.DB_HOST + ":" + strconv.Itoa(conf.Conf.DB_PORT) + ")/" + conf.Conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	maria, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	maria.AutoMigrate(&User{}, &Timestamp{})
}
