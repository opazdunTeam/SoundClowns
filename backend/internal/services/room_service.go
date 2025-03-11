package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type RoomService struct {
	roomRepo  repositories.IRoomRepository
	trackRepo repositories.ITrackRepository
	userRepo  repositories.IUserRepository
}

func NewRoomService(roomRepo repositories.IRoomRepository, trackRepo repositories.ITrackRepository, userRepo repositories.IUserRepository) *RoomService {
	return &RoomService{roomRepo: roomRepo, trackRepo: trackRepo, userRepo: userRepo}
}

// CreateRoom создаёт новую комнату.
func (s *RoomService) CreateRoom(name string, password string, ownerID uuid.UUID) (*models.Room, error) {
	owner, err := s.userRepo.FindByID(ownerID)
	if err != nil {
		return nil, errors.New("owner not found")
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	room := &models.Room{
		ID:       newUUID,
		Name:     name,
		Password: password,
		OwnerID:  owner.ID,
	}

	if err := s.roomRepo.Create(room); err != nil {
		return nil, err
	}

	return room, nil
}

// GetRooms возвращает список всех комнат.
func (s *RoomService) GetRooms() ([]*models.Room, error) {
	return s.roomRepo.GetRooms()
}

// GetTracks возвращает треки в комнате.
func (s *RoomService) GetTracks(roomID uuid.UUID) ([]*models.Track, error) {
	return s.roomRepo.GetTracks(roomID)
}

// FindByID находит комнату по её ID.
func (s *RoomService) FindByID(id string) (*models.Room, error) {
	return s.roomRepo.FindByID(id)
}

// FindByName находит комнату по её названию.
func (s *RoomService) FindByName(name string) (*models.Room, error) {
	return s.roomRepo.FindByName(name)
}

// UpdateTrackStatus обновляет статус трека.
func (s *RoomService) UpdateTrackStatus(trackID uuid.UUID, status string) error {
	switch status {
	case "approved":
		return s.trackRepo.ApproveTrack(trackID)
	case "rejected":
		return s.trackRepo.RejectTrack(trackID)
	default:
		return errors.New("invalid status")
	}
}

// DeleteRoom удаляет комнату.
func (s *RoomService) DeleteRoom(id uuid.UUID) error {
	return s.roomRepo.Delete(id.String())
}
