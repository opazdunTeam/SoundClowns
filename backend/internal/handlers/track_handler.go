package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type TrackHandler struct {
	trackService *services.TrackService
}

func NewTrackHandler(trackService *services.TrackService) *TrackHandler {
	return &TrackHandler{trackService: trackService}
}

// CreateTrack создает новый трек.
// @Summary Создать трек
// @Description Создает новый трек в указанной комнате.
// @Tags tracks
// @Accept json
// @Produce json
// @Param request body models.CreateTrackRequest true "Данные для создания трека"
// @Success 200 {object} models.Track
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks [post]
func (h *TrackHandler) CreateTrack(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RoomID    uuid.UUID      `json:"room_id"`
		UserID    uuid.UUID      `json:"user_id"`
		TrackData map[string]any `json:"track_data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	track := &models.Track{
		ID:        uuid.New(),
		RoomID:    req.RoomID,
		UserID:    req.UserID,
		TrackData: req.TrackData,
	}

	if err := h.trackService.Create(track); err != nil {
		http.Error(w, "Failed to create track", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}

// GetTrackByID возвращает трек по его ID.
// @Summary Получить трек по ID
// @Description Возвращает трек по его ID.
// @Tags tracks
// @Produce json
// @Param track_id path string true "ID трека"
// @Success 200 {object} models.Track
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks/{track_id} [get]
func (h *TrackHandler) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackIDStr := vars["track_id"]

	trackID, err := uuid.Parse(trackIDStr)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	track, err := h.trackService.FindByID(trackID)
	if err != nil {
		http.Error(w, "Track not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}

// UpdateTrack обновляет данные трека.
// @Summary Обновить трек
// @Description Обновляет данные трека по его ID.
// @Tags tracks
// @Accept json
// @Produce json
// @Param track_id path string true "ID трека"
// @Param request body models.Track true "Данные для обновления"
// @Success 200 {object} models.Track
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks/{track_id} [put]
func (h *TrackHandler) UpdateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackIDStr := vars["track_id"]

	trackID, err := uuid.Parse(trackIDStr)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	var track models.Track
	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Убедимся, что ID совпадает
	track.ID = trackID
	if err := h.trackService.Update(&track); err != nil {
		http.Error(w, "Failed to update track", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}

// DeleteTrack удаляет трек по ID.
// @Summary Удалить трек
// @Description Удаляет трек по его ID.
// @Tags tracks
// @Param track_id path string true "ID трека"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks/{track_id} [delete]
func (h *TrackHandler) DeleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackIDStr := vars["track_id"]

	trackID, err := uuid.Parse(trackIDStr)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	if err := h.trackService.Delete(trackID); err != nil {
		http.Error(w, "Failed to delete track", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ApproveTrack одобряет трек.
// @Summary Одобрить трек
// @Description Одобряет трек по его ID.
// @Tags tracks
// @Param track_id path string true "ID трека"
// @Success 200
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks/{track_id}/approve [put]
func (h *TrackHandler) ApproveTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackIDStr := vars["track_id"]

	trackID, err := uuid.Parse(trackIDStr)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	if err := h.trackService.ApproveTrack(trackID); err != nil {
		http.Error(w, "Failed to approve track", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RejectTrack отклоняет трек.
// @Summary Отклонить трек
// @Description Отклоняет трек по его ID.
// @Tags tracks
// @Param track_id path string true "ID трека"
// @Success 200
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tracks/{track_id}/reject [put]
func (h *TrackHandler) RejectTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackIDStr := vars["track_id"]

	trackID, err := uuid.Parse(trackIDStr)
	if err != nil {
		http.Error(w, "Invalid track ID", http.StatusBadRequest)
		return
	}

	if err := h.trackService.RejectTrack(trackID); err != nil {
		http.Error(w, "Failed to reject track", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
