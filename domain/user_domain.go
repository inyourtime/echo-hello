package domain

import "echo-hello/model"

type UserRepository interface {
	Create(newUser *model.User) (user *model.User, err error)
	FindByID(id uint) (user model.User, err error)
	FindByEmail(email string) (user model.User, err error)
	FindAll() (users []model.User, err error)
	Update(id uint, user *model.User) (err error)
}
