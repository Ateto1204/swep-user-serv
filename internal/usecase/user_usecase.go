package usecase

import (
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
)

type UserUseCase interface {
	SaveUser(id, name string) (*entity.User, error)
	GetUser(id string) (*domain.User, error)
}

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repo,
	}
}

func (uc *userUseCase) SaveUser(userID, name string) (*entity.User, error) {
	t := time.Now()
	user, err := uc.repository.Save(userID, name, t)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) GetUser(id string) (*domain.User, error) {
	user, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
