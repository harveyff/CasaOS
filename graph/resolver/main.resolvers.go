package graph_resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graph_generated "github.com/IceWhaleTech/CasaOS/graph/generated"
)

// Pong is the resolver for the pong field.
func (r *mutationResolver) Pong(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graph_generated.MutationResolver implementation.
func (r *Resolver) Mutation() graph_generated.MutationResolver { return &mutationResolver{r} }

// Query returns graph_generated.QueryResolver implementation.
func (r *Resolver) Query() graph_generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
