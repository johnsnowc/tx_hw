package main

import (
	"fmt"
	"log"
	api "tx_hw/kitex_gen/api/transaction"
	"tx_hw/model"
)

func main() {
	err := model.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql success")
	}
	// 延迟关闭数据库
	defer model.DB.Close()

	// 模型关闭 自动迁移【把结构体和数据表进行对应】
	model.DB.AutoMigrate(&model.Transaction{})

	svr := api.NewServer(new(TransactionImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
