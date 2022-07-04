package graph_resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graph_generated "github.com/IceWhaleTech/CasaOS/graph/generated"
	graph_generated_model "github.com/IceWhaleTech/CasaOS/graph/generated/model"
)

// AppList is the resolver for the appList field.
func (r *queryResolver) AppList(ctx context.Context, input graph_generated_model.AppListInput) (*graph_generated_model.AppListOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPort is the resolver for the getPort field.
func (r *queryResolver) GetPort(ctx context.Context, input graph_generated_model.PortInput) (*graph_generated_model.PortOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns graph_generated.QueryResolver implementation.
func (r *Resolver) Query() graph_generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
