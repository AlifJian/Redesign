package models

type Result struct {
	Status  int16
	Message string
	Data    []Event
}
