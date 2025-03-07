package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{1, "peter", "man", "swimming"}
	db.Create(&u1)

	var u UserInfo
	db.First(&u)
	fmt.Printf("%#v\n", u)

	db.Model(&u).Update("hobby", "programming")
	fmt.Printf("%#v\n", u)

	db.Delete(&u)
	fmt.Printf("%#v\n", u)
}
