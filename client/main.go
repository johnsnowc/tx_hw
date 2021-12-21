// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
	"tx_hw/kitex_gen/api"
	"tx_hw/kitex_gen/api/transaction"
	"tx_hw/model"
)

var DB *gorm.DB

func main() {
	client, err := transaction.NewClient("transaction", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	//add
	r.POST("tx", func(context *gin.Context) {
		var tx model.Transaction
		context.BindJSON(&tx)
		tx.CreateTime = time.Now()
		req := &api.AddRequest{
			Id:         tx.ID,
			RoomId:     tx.RoomID,
			ItemId:     tx.ItemID,
			Payment:    tx.Payment,
			CreateTime: tx.CreateTime.Format("2006-01-02 15:04:05"),
			UserId:     tx.UserID,
		}
		log.Println(tx.CreateTime)
		log.Println(tx.CreateTime.String())
		resp, err := client.Add(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, resp)
		}
	})

	//delete by id
	r.DELETE("tx/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &api.DeleteByIDRequest{id}
		resp, err := client.DeleteByID(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, resp)
		}
	})

	//select sum of payment by room_id
	r.GET("payments/:room_id", func(context *gin.Context) {
		RoomId, err := strconv.ParseInt(context.Param("room_id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &api.PaymentsByRoomIDRequest{RoomId}
		resp, err := client.PaymentsByRoomID(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, resp)
		}
	})

	r.Run()
}
