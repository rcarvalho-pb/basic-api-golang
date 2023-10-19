package repositories

import (
	"api/db/database"
	"context"
	"fmt"
)

type users struct {
	queries *database.Queries
}

func NewUserRepository(queries *database.Queries) *users {
	return &users{queries}
}

func (u users) CreateUser(user database.User) (uint32, error) {
	ctx := context.Background()

	result, err := u.queries.CreateUser(ctx, database.CreateUserParams{
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return 0, err
	}
	userIndex, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(userIndex), nil
}

func (u users) FindUser(nameOrNick string) ([]database.User, error) {
	ctx := context.Background()

	users, err := u.queries.FindUser(ctx, database.FindUserParams{
		Name: fmt.Sprintf("%%%s%%", nameOrNick),
		Nick: fmt.Sprintf("%%%s%%", nameOrNick),
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u users) GetUserById(userId int32) (database.User, error) {
	ctx := context.Background()

	user, err := u.queries.GetUserById(ctx, userId)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (u users) UpdateUserById(userId uint32, user database.User) error {
	ctx := context.Background()

	err := u.queries.UpdateUserById(ctx, database.UpdateUserByIdParams{
		ID:       int32(userId),
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u users) DeleteUserById(userId uint32) error {
	ctx := context.Background()

	if err := u.queries.DeleteUserById(ctx, int32(userId)); err != nil {
		return err
	}

	return nil
}

func (u users) GetUserByEmailOrNick(emailOrNick string) (database.User, error) {
	ctx := context.Background()

	user, err := u.queries.GetUserByEmailOrNick(ctx, database.GetUserByEmailOrNickParams{
		Email: emailOrNick,
		Nick:  emailOrNick,
	})
	if err != nil {
		return database.User{}, nil
	}

	return user, nil
}

func (u users) FollowUser(id, followedId uint32) error {
	ctx := context.Background()
	if _, err := u.queries.FollowUser(ctx, database.FollowUserParams{
		UserID:     int32(id),
		FollowerID: int32(followedId),
	}); err != nil {
		return err
	}

	return nil
}
