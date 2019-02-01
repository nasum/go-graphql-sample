package go_graphql_sample

import (
	"context"
	"fmt"
	"math/rand"
)

type Resolver struct {
	ships []Ship
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateShip(ctx context.Context, input NewShip) (Ship, error) {
	ship := Ship{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		ShipType: input.ShipType,
	}
	r.ships = append(r.ships, ship)
	return ship, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ships(ctx context.Context, name *string, shipType *ShipType) ([]Ship, error) {
	return r.ships, nil
}
