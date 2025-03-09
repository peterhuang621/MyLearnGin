package main

import (
	"gin_demo/demo_25/dao"
	"gin_demo/demo_25/models"
	"gin_demo/demo_25/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()

	r.Run(":8080")
}
