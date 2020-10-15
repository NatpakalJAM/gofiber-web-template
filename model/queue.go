package model

//QueueMessage queue message
type QueueMessage struct {
	ID   uint64      `json:"id"`
	Data interface{} `json:"data"`
}
