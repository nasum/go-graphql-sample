package go_graphql_sample

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateShip(ctx context.Context, input NewShip) (Ship, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ships(ctx context.Context, name *string, shipType *ShipType) ([]Ship, error) {
	panic("not implemented")
}
