package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"github.com/google/uuid"
)

type TrackService struct {
	trackRepo repositories.ITrackRepository
}

func NewTrackService(trackRepo repositories.ITrackRepository) *TrackService {
	return &TrackService{trackRepo: trackRepo}
}

// Create создаёт новый трек.
func (s *TrackService) Create(track *models.Track) error {
	return s.trackRepo.Create(track)
}

// FindByID находит трек по его ID.
func (s *TrackService) FindByID(id uuid.UUID) (*models.Track, error) {
	return s.trackRepo.FindByID(id)
}

// Update обновляет данные трека.
func (s *TrackService) Update(track *models.Track) error {
	return s.trackRepo.Update(track)
}

// Delete удаляет трек.
func (s *TrackService) Delete(id uuid.UUID) error {
	return s.trackRepo.Delete(id)
}

// ApproveTrack одобряет трек (меняет его статус).
func (s *TrackService) ApproveTrack(id uuid.UUID) error {
	return s.trackRepo.ApproveTrack(id)
}

// RejectTrack отклоняет трек (меняет его статус).
func (s *TrackService) RejectTrack(id uuid.UUID) error {
	return s.trackRepo.RejectTrack(id)
}
