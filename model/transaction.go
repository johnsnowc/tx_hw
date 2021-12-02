package model

import "time"

type Transaction struct {
	ID         int64     `json:"id"`
	RoomID     int64     `json:"room_id"`
	ItemID     int64     `json:"item_id"`
	Payment    float64   `json:"payment"`
	CreateTime time.Time `json:"create_time"`
	UserID     int64     `json:"user_id"`
}
