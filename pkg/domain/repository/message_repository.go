package repository

import (
	"github.com/hirosi1900day/go-graphql-server/pkg/domain/model"
)

type Message interface {
	GetMessages() ([]*model.Message, error)
	CreateMessage(todos *model.Message) (*model.Message, error)
}
