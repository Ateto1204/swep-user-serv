package repository

import (
	"encoding/json"
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(userID, name string, t time.Time) (*entity.User, error)
	GetByID(id string) (*domain.User, error)
	UpdFriends(user *domain.User) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(userID, name string, t time.Time) (*entity.User, error) {
	user := &entity.User{
		ID:       userID,
		Name:     name,
		Chats:    "[]",
		Friends:  "[]",
		CreateAt: t,
	}
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(userID string) (*domain.User, error) {
	var userEntity entity.User
	if err := r.db.Where("id = ?", userID).Order("id").First(&userEntity).Error; err != nil {
		return nil, err
	}
	chatsData, err := strUnserialize(userEntity.Chats)
	if err != nil {
		return nil, err
	}
	friendsData, err := strUnserialize(userEntity.Friends)
	if err != nil {
		return nil, err
	}
	userModel := &domain.User{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Chats:    chatsData,
		Friends:  friendsData,
		CreateAt: userEntity.CreateAt,
	}
	return userModel, err
}

func (r *userRepository) UpdFriends(user *domain.User) (*domain.User, error) {
	friendsData, err := strSerialize(user.Friends)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Friends:  friendsData,
		CreateAt: user.CreateAt,
	}
	if err := r.db.Model(userEntity).Update("Friends", friendsData).Error; err != nil {
		return nil, err
	}
	return r.GetByID(user.ID)
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
