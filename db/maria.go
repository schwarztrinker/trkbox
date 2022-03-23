package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var UsersDB Users
var maria *gorm.DB

func InitMariaDB() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	maria, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	maria.AutoMigrate(&Timestamp{}, &User{})

	//var new User
	//Maria.First(&new, 1)

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// result := Maria.Create(&user)

	// fmt.Println("#")
	// fmt.Println(user.ID)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
}
