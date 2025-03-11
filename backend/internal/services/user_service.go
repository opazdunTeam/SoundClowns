package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/pkg/yandexmusic"
	"fmt"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(token string) (*models.User, error) {
	client := yandexmusic.NewYandexMusicClient(token)

	account, err := client.AccountStatus()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account status: %w", err)
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	user := &models.User{
		ID:    newUUID,
		Token: token,
		UID:   account.UID,
		Login: account.Login,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]*models.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetUserByToken(token string) (*models.User, error) {
	return s.userRepo.FindByToken(token)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}
