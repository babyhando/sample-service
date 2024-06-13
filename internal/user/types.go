package user

import "context"

type Repo interface {
	Create(ctx context.Context, user *User) error
}

type User struct {
}
