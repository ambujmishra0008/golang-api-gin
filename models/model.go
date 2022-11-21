package models

import (
	"time"
)

type Request struct {
	Amount    string    `json:"“amount”,omitempty"`
	Timestamp time.Time `json:"“timestamp”,omitempty"`
}

type CorrReq struct {
	Amount    float64   `json:"“amount”,omitempty"`
	Timestamp time.Time `json:"“timestamp”,omitempty"`
}

type Response struct {
	Status  int     `json:"Status,omitempty"`
	Message string  `json:"Message,omitempty"`
	TransId string  `json:"TransId,omitempty"`
	Sum     float64 `json:"sum,omitempty"`
	Avg     float64 `json:"avg,omitempty"`
	Max     float64 `json:"max,omitempty"`
	Min     float64 `json:"min,omitempty"`
	Count   int     `json:"count,omitempty"`
}
