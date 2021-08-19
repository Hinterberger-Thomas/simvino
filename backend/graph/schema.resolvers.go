package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"simvino/auth"
	"simvino/graph/generated"
	"simvino/graph/model"
	"simvino/models/balances"
	"simvino/models/users"
)

func (r *mutationResolver) UpdateBalance(ctx context.Context, input model.NewTransaction) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("access denied")
	}
	balances.InsertBalance(&balances.Balance{UserID: user.UserID, Currency: input.Currency, Value: int(input.Value)})
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	err := user.InsertUser()

	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) GetBalance(ctx context.Context) ([]*model.Transaction, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}

	return balances.GetTransactionByUserID(user.UserID), nil
}

func (r *queryResolver) RefreshToken(ctx context.Context) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	return token, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
