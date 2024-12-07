package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(userID, name string, t time.Time) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
	UpdByID(field string, user *domain.User) (*domain.User, error)
	DeleteByID(userID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(userID, name string, t time.Time) (*domain.User, error) {
	userModel := domain.NewUser(userID, name, t)
	userEntity, err := parseToEntity(userModel)
	if err != nil {
		return nil, err
	}
	if err := r.db.Create(userEntity).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *userRepository) GetByID(userID string) (*domain.User, error) {
	var userEntity *entity.User
	if err := r.db.Where("id = ?", userID).Order("id").First(&userEntity).Error; err != nil {
		return nil, err
	}
	userModel, err := parseToModel(userEntity)
	if err != nil {
		userEntity.Notifs = "[]"
		userModel, err = parseToModel(userEntity)
		if err != nil {
			return nil, err
		}
		field := "Notifs"
		return r.UpdByID(field, userModel)
	}
	return userModel, err
}

func (r *userRepository) UpdByID(field string, user *domain.User) (*domain.User, error) {
	userEntity, err := parseToEntity(user)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(userEntity).Elem()
	f := v.FieldByName(field)
	if !f.IsValid() {
		return nil, errors.New("specified field does not exist in user entity")
	}

	if err := r.db.Model(userEntity).Update(field, f.Interface()).Error; err != nil {
		return nil, err
	}
	return r.GetByID(user.ID)
}

func (r *userRepository) DeleteByID(userID string) error {
	result := r.db.Where("id = ?", userID).Delete(&entity.User{})
	if result.Error != nil {
		return fmt.Errorf("error occur when deleting the user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user %s was not found", userID)
	}

	return nil
}

func parseToEntity(user *domain.User) (*entity.User, error) {
	chatsStr, err := strSerialize(user.Chats)
	if err != nil {
		return nil, err
	}
	friendsStr, err := strSerialize(user.Friends)
	if err != nil {
		return nil, err
	}
	notifsStr, err := strSerialize(user.Notifs)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		ID:       user.ID,
		Profile:  user.Profile,
		Name:     user.Name,
		Chats:    chatsStr,
		Friends:  friendsStr,
		Notifs:   notifsStr,
		CreateAt: user.CreateAt,
	}
	return userEntity, nil
}

func parseToModel(user *entity.User) (*domain.User, error) {
	chatsData, err := strUnserialize(user.Chats)
	if err != nil {
		return nil, err
	}
	friendsData, err := strUnserialize(user.Friends)
	if err != nil {
		return nil, err
	}
	notifsData, err := strUnserialize(user.Notifs)
	if err != nil {
		return nil, err
	}
	userModel := &domain.User{
		ID:       user.ID,
		Profile:  user.Profile,
		Name:     user.Name,
		Chats:    chatsData,
		Friends:  friendsData,
		Notifs:   notifsData,
		CreateAt: user.CreateAt,
	}
	return userModel, nil
}

func strSerialize(sa []string) (string, error) {
	s, err := json.Marshal(sa)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func strUnserialize(s string) ([]string, error) {
	var sa []string
	err := json.Unmarshal([]byte(s), &sa)
	return sa, err
}
