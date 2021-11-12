package user

import "github.com/startup-of-zero-reais/takeout-users-api/application/domains/user"

type User struct {
}

func NewUserService() *User {
	return &User{}
}

func (us *User) Create(u *user.User) (*user.User, error) {
	return u, nil
}

func (us *User) Update(id string, u *user.User) (*user.User, error) {
	return u, nil
}

func (us *User) GetOne(id string) (*user.User, error) {
	return new(user.User), nil
}

func (us *User) List(query string) ([]*user.User, error) {
	var users []*user.User

	users = append(users, new(user.User))

	return users, nil
}

func (us *User) Delete(id string) (*user.User, error) {
	return new(user.User), nil
}
