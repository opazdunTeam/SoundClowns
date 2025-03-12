package handlers

import (
	"backend/internal/services"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type RoomHandler struct {
	roomService *services.RoomService
}

func NewRoomHandler(roomService *services.RoomService) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}

// CreateRoom создает новую комнату.
// @Summary Создать комнату
// @Description Создает новую комнату с указанным именем, паролем и владельцем.
// @Tags rooms
// @Accept json
// @Produce json
// @Param request body models.CreateRoomRequest true "Данные для создания комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms [post]
func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string    `json:"name"`
		Password string    `json:"password"`
		OwnerID  uuid.UUID `json:"owner_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.CreateRoom(req.Name, req.Password, req.OwnerID)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// GetRooms возвращает список всех комнат.
// @Summary Получить все комнаты
// @Description Возвращает список всех комнат.
// @Tags rooms
// @Produce json
// @Success 200 {array} models.Room
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms [get]
func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.roomService.GetRooms()
	if err != nil {
		http.Error(w, "Failed to fetch rooms", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

// GetRoomByID возвращает комнату по ID.
// @Summary Получить комнату по ID
// @Description Возвращает комнату по её ID.
// @Tags rooms
// @Produce json
// @Param id path string true "ID комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{id} [get]
func (h *RoomHandler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.FindByID(id)
	if err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// GetRoomByName возвращает комнату по названию.
// @Summary Получить комнату по названию
// @Description Возвращает комнату по её названию.
// @Tags rooms
// @Produce json
// @Param name path string true "Название комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/name/{name} [get]
func (h *RoomHandler) GetRoomByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	room, err := h.roomService.FindByName(name)
	if err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// GetTracks возвращает треки в комнате.
// @Summary Получить треки в комнате
// @Description Возвращает список треков в указанной комнате.
// @Tags rooms
// @Produce json
// @Param room_id path string true "ID комнаты"
// @Success 200 {array} models.Track
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{room_id}/tracks [get]
func (h *RoomHandler) GetTracks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomIDStr := vars["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	tracks, err := h.roomService.GetTracks(roomID)
	if err != nil {
		http.Error(w, "Failed to fetch tracks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}

// GetTracksByStatus возвращает треки в комнате по статусу.
// @Summary Получить треки в комнате по статусу
// @Description Возвращает список треков в указанной комнате с заданным статусом.
// @Tags rooms
// @Produce json
// @Param room_id path string true "ID комнаты"
// @Param status query string true "Статус трека (pending, approved, rejected)"
// @Success 200 {array} models.Track
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{room_id}/tracks [get]
func (h *RoomHandler) GetTracksByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomIDStr := vars["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	status := r.URL.Query().Get("status")
	if status == "" {
		http.Error(w, "Status is required", http.StatusBadRequest)
		return
	}

	tracks, err := h.roomService.GetTracksByStatus(roomID, status)
	if err != nil {
		http.Error(w, "Failed to fetch tracks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}

// DeleteRoom удаляет комнату по ID.
// @Summary Удалить комнату
// @Description Удаляет комнату по её ID.
// @Tags rooms
// @Param id path string true "ID комнаты"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{id} [delete]
func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	if err := h.roomService.DeleteRoom(id); err != nil {
		http.Error(w, "Failed to delete room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// JoinRoom позволяет пользователю войти в комнату.
// @Summary Войти в комнату
// @Description Позволяет пользователю войти в комнату по логину и паролю.
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "ID комнаты"
// @Param request body models.JoinRoomRequest true "Данные для входа"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{id}/join [post]
func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	roomID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.JoinRoom(roomID, req.Password)
	if err != nil {
		http.Error(w, "Failed to join room", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// JoinRoomByLink позволяет пользователю войти в комнату по реферальной ссылке.
// @Summary Войти в комнату по ссылке
// @Description Позволяет пользователю войти в комнату по реферальной ссылке.
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "ID комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{id}/join-by-link [post]
func (h *RoomHandler) JoinRoomByLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	roomID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.JoinRoomByLink(roomID)
	if err != nil {
		http.Error(w, "Failed to join room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// LeaveRoom позволяет пользователю выйти из комнаты.
// @Summary Выйти из комнаты
// @Description Позволяет пользователю выйти из комнаты.
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "ID комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /rooms/{id}/leave [post]
func (h *RoomHandler) LeaveRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	roomID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	room, err := h.roomService.LeaveRoom(roomID)
	if err != nil {
		http.Error(w, "Failed to leave room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
