package entity

import "time"

type Gallery struct {
	Id string `json:"id"`
	Link string `json:"link"`
	UserId string `json:"userId"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}