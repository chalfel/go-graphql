package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/chalfel/go-graphql/database"
	"github.com/chalfel/go-graphql/graph/generated"
	"github.com/chalfel/go-graphql/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	res := db.Save(input)
	return res, nil
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	res := db.FindById(id)
	return res, nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	res := db.All()
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
