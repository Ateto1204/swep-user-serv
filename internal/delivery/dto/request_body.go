package dto

type SaveUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccessUserRequest struct {
	ID string `json:"id"`
}

type AccessFriendRequest struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
}

type AccessChatRequest struct {
	UserID string `json:"user_id"`
	ChatID string `json:"chat_id"`
}

type AccessNotifRequest struct {
	UserID  string `json:"user_id"`
	NotifID string `json:"notif_id"`
}

type AccessProfileRequest struct {
	UserID  string `json:"user_id"`
	Profile string `json:"profile"`
}

type AccessSettingRequest struct {
	UserID    string `json:"user_id"`
	Alias     string `json:"alias"`
	Birth     string `json:"birth"`
	Gender    string `json:"gen"`
	Telephone string `json:"tel"`
	Address   string `json:"addr"`
	Text      string `json:"txt"`
}
