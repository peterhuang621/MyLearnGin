package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'huang'"`
	Age  int64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	// u := User{Name: "peter", Age: 18}
	u := User{Age: 20}
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("User exist!")
	} else {
		fmt.Println("User not exist!")
	}
	db.Debug().Create(&u)
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("User exist!")
	} else {
		fmt.Println("User not exist!")
	}
}
