# Инструкция по запуску backend
## Перейти в папку backend
```bash
cd backend
```
## Создать и запустить контейнер с базой данных
```bash 
cd database
```
```bash
docker-compose up -d
```
```bash
cd ..
```
## Установить зависимости
```bash
go mod tidy
```
## Запустить api
```bash
go run cmd/api/main.go
```
