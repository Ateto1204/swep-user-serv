package repository

import (
	"github.com/Ateto1204/swep-user-serv/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User) error
	GetByID(id string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id string) (entity.User, error) {
	var user entity.User
	// err := r.db.First(&user, id).Error
	err := r.db.Where("id = ?", id).Order("id").First(&user).Error
	return user, err
}
