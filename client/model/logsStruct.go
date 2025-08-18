package model

import "time"

type LogsObject struct {
	StructTimeStamp time.Time `json:"timestamp"`
	StructMessage   string    `json:"message"`
}
