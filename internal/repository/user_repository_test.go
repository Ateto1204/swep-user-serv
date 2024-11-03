package repository_test

import (
	"testing"
	"time"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/Ateto1204/swep-user-serv/internal/domain"
	"github.com/Ateto1204/swep-user-serv/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDB() {
	// 使用 SQLite 內存資料庫進行測試，避免與生產資料庫衝突
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	// 自動建立資料表，確保測試表結構一致
	db.AutoMigrate(&entity.User{})
	testDB = db
}

func TestSave(t *testing.T) {
	setupTestDB()
	repo := repository.NewUserRepository(testDB)

	userID := "user123"
	name := "Test User"
	now := time.Now()

	user, err := repo.Save(userID, name, now)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, "[]", user.Chats)
	assert.Equal(t, "[]", user.Friends)
	assert.Equal(t, now, user.CreateAt)
}

func TestGetByID(t *testing.T) {
	setupTestDB()
	repo := repository.NewUserRepository(testDB)

	userID := "user123"
	name := "Test User"
	now := time.Now()
	repo.Save(userID, name, now)

	user, err := repo.GetByID(userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, []string{}, user.Chats)
	assert.Equal(t, []string{}, user.Friends)

	assert.True(t, user.CreateAt.Equal(now), "CreateAt should match")
}

func TestUpdFriends(t *testing.T) {
	setupTestDB()
	repo := repository.NewUserRepository(testDB)

	userID := "user123"
	name := "Test User"
	now := time.Now()
	user, _ := repo.Save(userID, name, now)

	userModel := &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Friends:  []string{"friend1", "friend2"},
		CreateAt: user.CreateAt,
	}
	updatedUser, err := repo.UpdFriends(userModel)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, userID, updatedUser.ID)
	assert.Equal(t, name, updatedUser.Name)
	assert.Equal(t, []string{"friend1", "friend2"}, updatedUser.Friends)
}
