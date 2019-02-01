package go_graphql_sample

import (
	"context"
	"fmt"
	"math/rand"
)

type Resolver struct {
	organizations []Organization
	ships         []Ship
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
func (r *mutationResolver) CreateOrganization(ctx context.Context, input NewOrganization) (Organization, error) {
	ships := []*Ship{}

	for _, ship := range input.Ships {
		newShip, _ := r.CreateShip(ctx, *ship)
		ships = append(ships, &newShip)
	}
	organization := Organization{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Name:  input.Name,
		Ships: ships,
	}
	r.organizations = append(r.organizations, organization)

	return organization, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ships(ctx context.Context, name *string, shipType *ShipType) ([]Ship, error) {
	return r.ships, nil
}
func (r *queryResolver) Organizations(ctx context.Context, name *string) ([]Organization, error) {
	return r.organizations, nil
}
