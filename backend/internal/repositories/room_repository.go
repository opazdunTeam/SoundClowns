package repositories

import (
	"backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) IRoomRepository { // Возвращаем интерфейс
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Create(room *models.Room) error {
	return r.db.Create(room).Error
}

func (r *RoomRepository) GetRooms() ([]*models.Room, error) {
	var rooms []*models.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) GetTracks(roomID uuid.UUID) ([]*models.Track, error) {
	var tracks []*models.Track
	if err := r.db.Where("room_id = ?", roomID).Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

func (r *RoomRepository) FindByID(id string) (*models.Room, error) {
	var room models.Room
	if err := r.db.Where("id = ?", id).First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) FindByName(name string) (*models.Room, error) {
	var room models.Room
	if err := r.db.Where("name = ?", name).First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) Update(room *models.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Room{}).Error
}
