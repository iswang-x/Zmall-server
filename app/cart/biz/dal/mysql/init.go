package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
