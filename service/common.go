package service

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

const(
	InvalidRequest = "request is invalid"
	usernameOrEmailIsDuplicate = "username or Email is duplicate"
)

type Handler struct{
	db *gorm.DB
	p *kafka.Producer
}

func NewHandler(db *gorm.DB, p *kafka.Producer) *Handler {
	return &Handler{
		db: db,
		p: p}
}
