package database

import (
	"ders-programi/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {
	Database()
	Migrate()
}

func Database() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang-ders-programi?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database bağlantısı başarılı")
}

func Migrate() error {
	err := Conn.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	fmt.Println("Migration başarılı")

	return nil
}
