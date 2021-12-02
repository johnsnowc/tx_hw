package main

import (
	"context"
	"log"
	"time"
	"tx_demo/kitex_gen/api"
	"tx_demo/model"
)

// TransactionImpl implements the last service interface defined in the IDL.
type TransactionImpl struct{}

// Add implements the TransactionImpl interface.
func (s *TransactionImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", req.CreateTime, time.Local)
	//time:=time.Now()
	tx := &model.Transaction{
		req.Id,
		req.RoomId,
		req.ItemId,
		req.Payment,
		time,
		req.UserId,
	}
	err = DB.Create(&tx).Error
	if err != nil {
		log.Fatal(err)
		resp = &api.AddResponse{"add failed!", req.Id}
		return
	}
	resp = &api.AddResponse{"add succeed!", req.Id}
	return
}

// DeleteByID implements the TransactionImpl interface.
func (s *TransactionImpl) DeleteByID(ctx context.Context, req *api.DeleteByIDRequest) (resp *api.DeleteByIDResponse, err error) {
	// TODO: Your code here...
	id := req.Id
	err = DB.Where("id = ?", id).Delete(model.Transaction{}).Error
	if err != nil {
		log.Fatal(err)
		resp = &api.DeleteByIDResponse{"delete failed!", req.Id}
		return
	}
	resp = &api.DeleteByIDResponse{"delete succeed!", req.Id}
	return
}

// PaymentsByRoomID implements the TransactionImpl interface.
func (s *TransactionImpl) PaymentsByRoomID(ctx context.Context, req *api.PaymentsByRoomIDRequest) (resp *api.PaymentsByRoomIDResponse, err error) {
	// TODO: Your code here...
	RoomId := req.RoomId
	var txs []model.Transaction
	err = DB.Debug().Where("room_id = ?", RoomId).Find(&txs).Error
	if err != nil {
		log.Fatal(err)
		resp = &api.PaymentsByRoomIDResponse{0}
		return
	}
	var sum float64
	for _, tx := range txs {
		sum += tx.Payment
	}
	resp = &api.PaymentsByRoomIDResponse{sum}
	return
}
