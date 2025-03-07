package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	var u User

	db.Debug().Where("name=?", "xiaowangzi").Delete(&u)
	fmt.Printf("%#v\n", u)

	db.Debug().Unscoped().Where("name=?", "huang").Delete(User{})
}
