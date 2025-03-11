package main

import (
	_ "backend/docs"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {
	// Подключение к базе данных
	db, err := repositories.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	roomRepo := repositories.NewRoomRepository(db)
	trackRepo := repositories.NewTrackRepository(db)

	// Инициализация сервисов
	userService := services.NewUserService(userRepo)
	roomService := services.NewRoomService(roomRepo, trackRepo, userRepo)
	trackService := services.NewTrackService(trackRepo)

	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(userService)
	roomHandler := handlers.NewRoomHandler(roomService)
	trackHandler := handlers.NewTrackHandler(trackService)

	// Настройка маршрутов
	r := mux.NewRouter()

	// Swagger UI
	/*
		go install github.com/swaggo/swag/cmd/swag@latest
		swag init -g cmd/api/main.go
		http://localhost:8080/swagger/
	*/
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Пользователи
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/token", userHandler.GetUserByToken).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Комнаты
	r.HandleFunc("/rooms", roomHandler.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms/{id}", roomHandler.GetRoomByID).Methods("GET")
	r.HandleFunc("/rooms/name/{name}", roomHandler.GetRoomByName).Methods("GET")
	r.HandleFunc("/rooms/{room_id}/tracks", roomHandler.GetTracks).Methods("GET")
	r.HandleFunc("/rooms", roomHandler.GetRooms).Methods("GET")
	r.HandleFunc("/rooms/{id}", roomHandler.DeleteRoom).Methods("DELETE")

	// Треки
	r.HandleFunc("/tracks", trackHandler.CreateTrack).Methods("POST")
	r.HandleFunc("/tracks/{track_id}", trackHandler.GetTrackByID).Methods("GET")
	r.HandleFunc("/tracks/{track_id}", trackHandler.UpdateTrack).Methods("PUT")
	r.HandleFunc("/tracks/{track_id}", trackHandler.DeleteTrack).Methods("DELETE")
	r.HandleFunc("/tracks/{track_id}/approve", trackHandler.ApproveTrack).Methods("PUT")
	r.HandleFunc("/tracks/{track_id}/reject", trackHandler.RejectTrack).Methods("PUT")

	// Запуск сервера
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
