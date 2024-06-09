package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/hirosi1900day/go-graphql-server/pkg/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MsgUseCase  usecase.Message
	UserUseCase usecase.User
}
