.PHONY: start stop install-frontend install-backend install clean lint lint-frontend lint-backend

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

lint: lint-frontend lint-backend

lint-backend-fix:
	@echo "Fixing backend lint issues..."
	cd backend && golangci-lint run --fix ./...

lint-frontend-fix:
	@echo "Fixing frontend lint issues..."
	cd app && npm run lint -- --fix

lint-fix: lint-fix-frontend lint-fix-backend

lint-frontend:
	@echo "Linting frontend..."
	cd app && npm run lint

format:
	@echo "Formatting..."
	cd app && npm run format && cd ../backend && gofmt -w .

lint-backend:
	@echo "Linting backend..."
	cd backend && golangci-lint run

pre-commit: lint
	@echo "Running pre-commit hooks..."
	pre-commit run --all-files 