package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	api "tx_demo/kitex_gen/api/transaction"
	"tx_demo/model"
)

var DB *gorm.DB

func main() {
	err := model.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql success")
	}
	// 延迟关闭数据库
	defer DB.Close()

	// 模型关闭 自动迁移【把结构体和数据表进行对应】
	DB.AutoMigrate(&model.Transaction{})

	svr := api.NewServer(new(TransactionImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
