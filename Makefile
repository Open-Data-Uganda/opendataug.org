.PHONY: start stop install-frontend install-backend install clean

all: install start

install: install-frontend install-backend

install-frontend:
	cd app && npm install

install-backend:
	cd backend && go mod download

start:
	@echo "Starting both frontend and backend applications..."
	@make start-backend & make start-frontend

start-frontend:
	@echo "Starting frontend application..."
	cd app && npm run dev

start-backend:
	@echo "Starting backend application..."
	cd backend && air

stop:
	@echo "Stopping all processes..."
	@pkill -f "npm run dev" || true
	@pkill -f "go run main.go" || true

clean: stop
	@echo "Cleaning up..."
	@rm -rf app/node_modules
	@rm -rf backend/tmp 