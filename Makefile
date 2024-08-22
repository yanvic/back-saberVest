include .env

up:
	@echo "Starting container..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping containers..."
	docker-compose down

build:
	go build -o ${BINARY} ./main.go

start:
	./${BINARY}

restart: build start