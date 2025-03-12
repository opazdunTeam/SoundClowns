package repositories

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	// Create создает нового пользователя в базе данных.
	Create(user *models.User) error

	// GetUsers возвращает список всех пользователей.
	GetUsers() ([]*models.User, error)

	// FindByID находит пользователя по его ID.
	FindByID(id uuid.UUID) (*models.User, error)

	// FindByToken находит пользователя по его OAuth токену.
	FindByToken(token string) (*models.User, error)

	// Update обновляет данные пользователя.
	Update(user *models.User) error

	// Delete удаляет пользователя по его ID.
	Delete(id uuid.UUID) error
}

type IRoomRepository interface {
	// Create создаёт новую комнату в базе данных.
	Create(room *models.Room) error

	// GetRooms возвращает список всех комнат.
	GetRooms() ([]*models.Room, error)

	// GetTracks возвращает треки по id комнаты
	GetTracks(roomID uuid.UUID) ([]*models.Track, error)

	GetTracksByStatus(roomID uuid.UUID, status string) ([]*models.Track, error)

	//GetApprovedTracks(roomID uuid.UUID) ([]*models.Track, error)
	//
	//GetRejectedTracks(roomID uuid.UUID) ([]*models.Track, error)

	// FindByID находит комнату по её ID.
	FindByID(id uuid.UUID) (*models.Room, error)

	// FindByName находит комнату по её названию.
	FindByName(name string) (*models.Room, error)

	// Update обновляет данные комнаты.
	Update(room *models.Room) error

	// Delete удаляет комнату.
	Delete(id uuid.UUID) error

	// JoinRoom позволяет пользователю войти в комнату по паролю.
	JoinRoom(roomID uuid.UUID, password string) (*models.Room, error)

	// JoinRoomByLink позволяет пользователю войти в комнату по реферальной ссылке.
	JoinRoomByLink(roomID uuid.UUID) (*models.Room, error)

	// LeaveRoom уменьшает счетчик пользователей в комнате.
	LeaveRoom(roomID uuid.UUID) error

	//JoinRoom()
	//LeaveRoom()
	//GetTracks()
	//SuggestTrack()
}

type ITrackRepository interface {
	// Create создаёт новый трек в базе данных.
	Create(track *models.Track) error

	// FindByID находит трек по его ID.
	FindByID(id uuid.UUID) (*models.Track, error)

	// Update обновляет данные трека.
	Update(track *models.Track) error

	// Delete удаляет трек.
	Delete(id uuid.UUID) error

	// ApproveTrack одобряет трек (меняет его статус).
	ApproveTrack(id uuid.UUID) error

	// RejectTrack отклоняет трек (меняет его статус).
	RejectTrack(id uuid.UUID) error
}
