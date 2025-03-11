package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser создает нового пользователя.
// @Summary Создать пользователя
// @Description Создает нового пользователя на основе OAuth токена.
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Данные для создания пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(req.Token)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUsers возвращает список всех пользователей.
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей.
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID возвращает пользователя по ID.
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по его ID.
// @Tags users
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUserByToken возвращает пользователя по токену.
// @Summary Получить пользователя по токену
// @Description Возвращает пользователя по его OAuth токену.
// @Tags users
// @Produce json
// @Param Authorization header string true "OAuth токен"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/token [get]
func (h *UserHandler) GetUserByToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByToken(token)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser обновляет данные пользователя.
// @Summary Обновить пользователя
// @Description Обновляет данные пользователя по его ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param request body models.User true "Данные для обновления"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Убедимся, что ID совпадает
	user.ID = id
	if err := h.userService.UpdateUser(&user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser удаляет пользователя по ID.
// @Summary Удалить пользователя
// @Description Удаляет пользователя по его ID.
// @Tags users
// @Param id path string true "ID пользователя"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
