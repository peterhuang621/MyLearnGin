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
	db.Debug().Where(User{Name: "peter"}).Find(&u)
	fmt.Printf("user: %#v\n", u)
	db.Debug().Save(&u)
	db.Debug().Model(&u).Update("name", "xiaowangzi")

	m1 := map[string]interface{}{
		"name": "huangyi",
		"age":  28,
	}
	db.Debug().Model(&u).Select("age").Updates(&m1)
}
