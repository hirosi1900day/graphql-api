package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"

	"github.com/WebEngrChild/go-graphql-server/pkg/domain/model"
	"github.com/WebEngrChild/go-graphql-server/pkg/lib/graph/generated"
)

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context) ([]*model.Message, error) {
	todos, err := r.MsgUseCase.GetMessages()
	if err != nil {
		err = fmt.Errorf("resolver Todos() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return todos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
