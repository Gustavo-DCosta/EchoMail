package model

import "time"

type ErrorObject struct {
	StructTime  time.Time `json:"timestamp"`
	StructError string    `json:"error"`
}
