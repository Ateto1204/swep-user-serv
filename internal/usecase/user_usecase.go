package usecase

import (
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
)

type UserUseCase interface {
	SaveUser(id, name string) error
	GetUser(id string) (*entity.User, error)
	Run()
}

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repo,
	}
}

func (uc *userUseCase) SaveUser(id, name string) error {
	t := time.Now()
	user := &entity.User{
		ID:       id,
		Name:     name,
		CreateAt: t,
	}
	err := uc.repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) GetUser(id string) (*entity.User, error) {
	user, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *userUseCase) Run() {

}
