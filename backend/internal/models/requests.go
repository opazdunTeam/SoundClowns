package models

import "github.com/google/uuid"

type CreateUserRequest struct {
	Token string `json:"token"`
}

type CreateRoomRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	OwnerID  string `json:"owner_id"`
}

type JoinRoomRequest struct {
	Password string `json:"password" example:"your_password"`
}

type CreateTrackRequest struct {
	RoomID    uuid.UUID      `json:"room_id"`
	UserID    uuid.UUID      `json:"user_id"`
	TrackData map[string]any `json:"track_data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
