package graph_resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graph_generated_model "github.com/IceWhaleTech/CasaOS/graph/generated/model"
)

// SetNamePwd is the resolver for the setNamePwd field.
func (r *mutationResolver) SetNamePwd(ctx context.Context, username string, pwd string) (*graph_generated_model.SetNamePwdOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserLogin is the resolver for the userLogin field.
func (r *mutationResolver) UserLogin(ctx context.Context, username string, pwd string) (*graph_generated_model.UserLoginOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserInfo is the resolver for the getUserInfo field.
func (r *queryResolver) GetUserInfo(ctx context.Context, id string) (*graph_generated_model.GetUserInfoOutput, error) {
	panic(fmt.Errorf("not implemented"))
}
