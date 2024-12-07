package delivery

import (
	"net/http"

	"github.com/Ateto1204/swep-user-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (h *UserHandler) SaveUser(c *gin.Context) {
	type Input struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.SaveUser(input.ID, input.Name)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.GetUser(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userUseCase.DeleteUser(input.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete user successfully"})
}

func (h *UserHandler) AddNewFriend(c *gin.Context) {
	type Input struct {
		UserID   string `json:"user_id"`
		FriendID string `json:"friend_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.AddNewFriend(input.UserID, input.FriendID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveFriend(c *gin.Context) {
	type Input struct {
		UserID   string `json:"user_id"`
		FriendID string `json:"friend_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.RemoveFriend(input.UserID, input.FriendID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) AddNewChat(c *gin.Context) {
	type Input struct {
		UserID string `json:"user_id"`
		ChatID string `json:"chat_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.AddNewChat(input.UserID, input.ChatID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveChat(c *gin.Context) {
	type Input struct {
		UserID string `json:"user_id"`
		ChatID string `json:"chat_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.RemoveChat(input.UserID, input.ChatID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) AddNewNotif(c *gin.Context) {
	type Input struct {
		UserID  string `json:"user_id"`
		NotifID string `json:"notif_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.AddNewNotif(input.UserID, input.NotifID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveNotif(c *gin.Context) {
	type Input struct {
		UserID  string `json:"user_id"`
		NotifID string `json:"notif_id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.RemoveNotif(input.UserID, input.NotifID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdProfileUrl(c *gin.Context) {
	type Input struct {
		UserID  string `json:"user_id"`
		Profile string `json:"profile"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.UpdProfileUrl(input.UserID, input.Profile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, "hello zeabur")
}
