package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
)

type UserUseCase interface {
	SaveUser(id, name string) (*entity.User, error)
	GetUser(id string) (*domain.User, error)
	AddNewChat(userID, chatID string) (*domain.User, error)
	AddNewFriend(userID, friendID string) (*domain.User, error)
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

func (uc *userUseCase) GetUser(userID string) (*domain.User, error) {
	user, err := uc.repository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) AddNewChat(userID, chatID string) (*domain.User, error) {
	// workflow: check if chat exists
	return nil, nil
}

func (uc *userUseCase) AddNewFriend(userID, friendID string) (*domain.User, error) {
	// workflow: check if id are same --> check if user exists --> checkout if they are already friends --> add friend
	if userID == friendID {
		return nil, errors.New("disable to add yourself as a friend")
	}

	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}

	friend, err := uc.repository.GetByID(friendID)
	if friend == nil || err != nil {
		return nil, err
	}

	checkChan := make(chan error, 2)

	go func() {
		for _, id := range user.Friends {
			if id == friendID {
				checkChan <- fmt.Errorf("user %s and friend %s are already friends", userID, friendID)
				return
			}
		}
		checkChan <- nil
	}()

	go func() {
		for _, id := range friend.Friends {
			if id == userID {
				checkChan <- fmt.Errorf("user %s and friend %s are already friends", friendID, userID)
				return
			}
		}
		checkChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-checkChan; err != nil {
			return nil, err
		}
	}

	user.Friends = append(user.Friends, friendID)
	friend.Friends = append(friend.Friends, userID)
	if _, err := uc.repository.AddNewFriend(friend); err != nil {
		return nil, err
	}

	return uc.repository.AddNewFriend(user)
}

func (uc *userUseCase) RemoveFriend(userID, friendID string) error {
	return nil
}
