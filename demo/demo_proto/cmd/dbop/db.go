package main

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "asdfjkl"})
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("Password", "22222")
	// var raw model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&raw)
	// fmt.Printf("row: %+v\n", raw)
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}
