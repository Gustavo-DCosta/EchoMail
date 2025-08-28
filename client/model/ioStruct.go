package model

import "time"

type LogsObject struct {
	StructTimeStamp time.Time `json:"timestamp"`
	StructMessage   string    `json:"message"`
}

type ErrorObject struct {
	StructTime  time.Time `json:"timestamp"`
	StructError string    `json:"error"`
}

type EmailObject struct {
	StructEmailObject string `json:"saved_email"`
}

type SessionObject struct {
	StructEmail       string    `json:"email"`
	StructPhoneNumber string    `json:"phone_number"`
	StructTimePrompt  time.Time `json:"date"`
}
