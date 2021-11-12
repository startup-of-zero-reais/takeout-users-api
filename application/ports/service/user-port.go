package service

import "github.com/startup-of-zero-reais/takeout-users-api/application/domains/user"

type (
	User interface {
		Create(u *user.User) (*user.User, error)
		Update(id string, u *user.User) (*user.User, error)
		GetOne(id string) (*user.User, error)
		List(query string) ([]*user.User, error)
		Delete(id string) (*user.User, error)
	}
)
