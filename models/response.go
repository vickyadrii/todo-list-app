package models

type Response struct {
	Message string      `json:"message"`
	Status  uint        `json:"status"`
	Data    interface{} `json:"data"`
}
