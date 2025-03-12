package repositories

import (
	"backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) ITrackRepository { // Возвращаем интерфейс
	return &TrackRepository{db: db}
}

func (r *TrackRepository) Create(track *models.Track) error {
	return r.db.Create(track).Error
}

func (r *TrackRepository) FindByID(id uuid.UUID) (*models.Track, error) {
	var track models.Track
	if err := r.db.Where("id = ?", id).First(&track).Error; err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *TrackRepository) Update(track *models.Track) error {
	return r.db.Save(track).Error
}

func (r *TrackRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&models.Track{}).Error
}

func (r *TrackRepository) ApproveTrack(id uuid.UUID) error {
	return r.db.Model(&models.Track{}).Where("id = ?", id).Update("status", "approved").Error
}

func (r *TrackRepository) RejectTrack(id uuid.UUID) error {
	return r.db.Model(&models.Track{}).Where("id = ?", id).Update("status", "rejected").Error
}
