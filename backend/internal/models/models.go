package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Login        string    `gorm:"type:text" json:"login"`
	Token        string    `gorm:"type:text;not null" json:"token"`
	UID          int64     `gorm:"type:bigint" json:"uid"`
	RegisteredAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"registered_at"`
}

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name       string    `gorm:"type:text;unique;not null" json:"name"`
	Password   string    `gorm:"type:text;not null" json:"-"`
	OwnerID    uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Owner      User      `gorm:"foreignKey:OwnerID" json:"-"`
	UsersCount int64     `gorm:"-" json:"users_count"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Track struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomID    uuid.UUID      `gorm:"type:uuid;not null" json:"room_id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	TrackData map[string]any `gorm:"type:jsonb;not null" json:"track_data"`
	Status    string         `gorm:"type:text;default:'pending'" json:"status"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Room      Room           `gorm:"foreignKey:RoomID" json:"-"` // Связь с таблицей rooms
	User      User           `gorm:"foreignKey:UserID" json:"-"` // Связь с таблицей users
}
