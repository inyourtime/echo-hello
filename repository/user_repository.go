package repository

import (
	"echo-hello/domain"
	"echo-hello/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(newUser *model.User) (user *model.User, err error) {
	err = ur.db.Create(newUser).Error
	if err != nil {
		return
	}
	return newUser, nil
}

func (ur *userRepository) FindByID(id uint) (user model.User, err error) {
	err = ur.db.First(&user, id).Error
	return
}

func (ur *userRepository) FindByEmail(email string) (user model.User, err error) {
	err = ur.db.Where("email = ?", email).First(&user).Error
	return
}

func (ur *userRepository) FindAll() (users []model.User, err error) {
	err = ur.db.Find(&users).Error
	return
}

func (ur *userRepository) Update(id uint, user *model.User) (err error) {
	return nil
}
