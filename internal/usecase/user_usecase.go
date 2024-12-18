package usecase

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Ateto1204/swep-user-serv/internal/delivery/dto"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
)

type UserUseCase interface {
	SaveUser(id, name string) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	DeleteUser(userID string) error
	AddNewChat(userID, chatID string) (*domain.User, error)
	RemoveChat(userID, chatID string) (*domain.User, error)
	AddNewFriend(userID, friendID string) (*domain.User, error)
	RemoveFriend(userID, friendID string) (*domain.User, error)
	AddNewNotif(userID, notifID string) (*domain.User, error)
	RemoveNotif(userID, notifID string) (*domain.User, error)
	UpdProfileUrl(userID, profileUrl string) (*domain.User, error)
	SaveSettings(settings dto.AccessSettingRequest) (*domain.User, error)
}

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repo,
	}
}

func (uc *userUseCase) SaveUser(userID, name string) (*domain.User, error) {
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

func (uc *userUseCase) DeleteUser(userID string) error {
	if _, err := uc.repository.GetByID(userID); err != nil {
		return err
	}
	if err := uc.repository.DeleteByID(userID); err != nil {
		return err
	}
	t := time.Now()
	log.Printf("delete user %s at %v", userID, t)
	return nil
}

func (uc *userUseCase) AddNewChat(userID, chatID string) (*domain.User, error) {
	// workflow: check if user exists --> check if user does not have the chat --> add chat
	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}
	for _, id := range user.Chats {
		if id == chatID {
			return nil, fmt.Errorf("user %s already have the chat %s", userID, chatID)
		}
	}
	user.Chats = append(user.Chats, chatID)

	return uc.repository.UpdByID("Chats", user)
}

func (uc *userUseCase) RemoveChat(userID, chatID string) (*domain.User, error) {
	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}
	flag := true
	for _, id := range user.Chats {
		if id == chatID {
			flag = false
			break
		}
	}
	if flag {
		return nil, fmt.Errorf("user %s doesn't have the chat %s", userID, chatID)
	}
	user.Chats = removeFromSlice(user.Chats, chatID)
	return uc.repository.UpdByID("Chats", user)
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
	if _, err := uc.repository.UpdByID("Friends", friend); err != nil {
		return nil, err
	}

	return uc.repository.UpdByID("Friends", user)
}

func (uc *userUseCase) RemoveFriend(userID, friendID string) (*domain.User, error) {
	// workflow: check if id are same --> check if user exists --> checkout if they are even not friends --> remove friend
	if userID == friendID {
		return nil, errors.New("disable to remove yourself as a friend")
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
				checkChan <- nil
				return
			}
		}
		checkChan <- fmt.Errorf("user %s and friend %s are even not friends", userID, friendID)
	}()

	go func() {
		for _, id := range friend.Friends {
			if id == userID {
				checkChan <- nil
				return
			}
		}
		checkChan <- fmt.Errorf("user %s and friend %s are even not friends", friendID, userID)
	}()

	for i := 0; i < 2; i++ {
		if err := <-checkChan; err != nil {
			return nil, err
		}
	}

	user.Friends = removeFromSlice(user.Friends, friendID)
	friend.Friends = removeFromSlice(friend.Friends, userID)
	if _, err := uc.repository.UpdByID("Friends", friend); err != nil {
		return nil, err
	}

	return uc.repository.UpdByID("Friends", user)
}

func (uc *userUseCase) AddNewNotif(userID, notifID string) (*domain.User, error) {
	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}

	user.Notifs = append(user.Notifs, notifID)
	field := "Notifs"
	user, err = uc.repository.UpdByID(field, user)
	return user, err
}

func (uc *userUseCase) RemoveNotif(userID, notifID string) (*domain.User, error) {
	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}

	user.Notifs = removeFromSlice(user.Notifs, notifID)
	field := "Notifs"
	user, err = uc.repository.UpdByID(field, user)
	return user, err
}

func (uc *userUseCase) UpdProfileUrl(userID, profileUrl string) (*domain.User, error) {
	user, err := uc.repository.GetByID(userID)
	if user == nil || err != nil {
		return nil, err
	}

	user.Profile = profileUrl
	field := "Profile"
	user, err = uc.repository.UpdByID(field, user)
	return user, err
}

func (uc *userUseCase) SaveSettings(set dto.AccessSettingRequest) (*domain.User, error) {
	user, err := uc.repository.GetByID(set.UserID)
	if err != nil {
		return nil, err
	}
	newSettings := []string{set.Alias, set.Birth, set.Gender, set.Telephone, set.Address, set.Text}
	user.Settings = newSettings
	field := "Settings"
	user, err = uc.repository.UpdByID(field, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func removeFromSlice(slice []string, target string) []string {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
