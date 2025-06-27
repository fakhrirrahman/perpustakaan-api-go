# Makefile untuk memudahkan development

# Variables
APP_NAME=perpustakaan-api
GO_VERSION=1.24.1

# Default target
.DEFAULT_GOAL := help

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

## help: Menampilkan bantuan
help:
	@echo "$(BLUE)=== Perpustakaan API - Go $(NC)"
	@echo ""
	@echo "$(GREEN)Tersedia commands:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

## install: Install dependencies
install:
	@echo "$(BLUE)Installing dependencies...$(NC)"
	go mod download
	go mod tidy

## build: Build aplikasi
build:
	@echo "$(BLUE)Building application...$(NC)"
	go build -o bin/$(APP_NAME) main.go
	@echo "$(GREEN)✓ Build completed!$(NC)"

## run: Menjalankan aplikasi
run:
	@echo "$(BLUE)Starting application...$(NC)"
	go run main.go

## run-dev: Menjalankan dengan auto-reload (install air terlebih dahulu)
run-dev:
	@echo "$(BLUE)Starting development server with auto-reload...$(NC)"
	@if ! command -v air > /dev/null; then \
		echo "$(YELLOW)Air not found. Installing...$(NC)"; \
		go install github.com/air-verse/air@latest; \
	fi
	export PATH=$$PATH:$(shell go env GOPATH)/bin && air

## test: Menjalankan test
test:
	@echo "$(BLUE)Running tests...$(NC)"
	go test -v ./...

## migrate: Menjalankan database migration (auto migration GORM)
migrate:
	@echo "$(BLUE)Running database migration...$(NC)"
	@echo "Migration akan berjalan otomatis saat aplikasi start (GORM AutoMigrate)"
	@echo "$(GREEN)✓ Migration setup completed!$(NC)"

## db-reset: Reset database (hati-hati - akan menghapus semua data!)
db-reset:
	@echo "$(RED)WARNING: This will delete all data!$(NC)"
	@read -p "Are you sure? (y/N): " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		echo "$(BLUE)Resetting database...$(NC)"; \
		echo "Menjalankan MySQL reset commands:"; \
		mysql -u root -p1 -e "DROP DATABASE IF EXISTS belajargo; CREATE DATABASE belajargo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"; \
		echo "$(GREEN)✓ Database reset completed!$(NC)"; \
	else \
		echo "$(YELLOW)Database reset cancelled.$(NC)"; \
	fi

## clean: Membersihkan build files
clean:
	@echo "$(BLUE)Cleaning build files...$(NC)"
	rm -rf bin/
	go clean
	@echo "$(GREEN)✓ Clean completed!$(NC)"

## mod-update: Update semua dependencies
mod-update:
	@echo "$(BLUE)Updating dependencies...$(NC)"
	go get -u ./...
	go mod tidy
	@echo "$(GREEN)✓ Dependencies updated!$(NC)"

## lint: Menjalankan golangci-lint
lint:
	@echo "$(BLUE)Running linter...$(NC)"
	@if ! command -v golangci-lint > /dev/null; then \
		echo "$(YELLOW)golangci-lint not found. Install it first:$(NC)"; \
		echo "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	else \
		golangci-lint run; \
	fi

## format: Format code dengan gofmt dan goimports
format:
	@echo "$(BLUE)Formatting code...$(NC)"
	gofmt -s -w .
	@if command -v goimports > /dev/null; then \
		goimports -w .; \
	else \
		echo "$(YELLOW)goimports not found. Install it:$(NC)"; \
		echo "go install golang.org/x/tools/cmd/goimports@latest"; \
	fi
	@echo "$(GREEN)✓ Code formatted!$(NC)"

## setup: Setup development environment
setup: install
	@echo "$(BLUE)Setting up development environment...$(NC)"
	@echo "$(GREEN)Installing useful development tools...$(NC)"
	go install github.com/air-verse/air@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(GREEN)✓ Development environment ready!$(NC)"
	@echo ""
	@echo "$(BLUE)Next steps:$(NC)"
	@echo "1. Copy .env.example to .env dan sesuaikan database config"
	@echo "2. Pastikan MySQL sudah running"
	@echo "3. Jalankan: $(YELLOW)make run$(NC)"

## docker-build: Build Docker image
docker-build:
	@echo "$(BLUE)Building Docker image...$(NC)"
	docker build -t $(APP_NAME) .

## docker-run: Menjalankan aplikasi dalam Docker
docker-run:
	@echo "$(BLUE)Running application in Docker...$(NC)"
	docker run -p 8080:8080 $(APP_NAME)

## logs: Menampilkan logs aplikasi (jika running dengan Docker)
logs:
	docker logs -f $(APP_NAME)

## swagger: Generate Swagger documentation
swagger:
	@echo "$(BLUE)Generating Swagger documentation...$(NC)"
	@if ! command -v swag > /dev/null; then \
		echo "$(YELLOW)Swag not found. Installing...$(NC)"; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	export PATH=$$PATH:$(shell go env GOPATH)/bin && swag init
	@echo "$(GREEN)✓ Swagger documentation generated!$(NC)"
	@echo "Access documentation at: $(YELLOW)http://localhost:8000/swagger/index.html$(NC)"

## swagger-serve: Generate documentation dan serve aplikasi
swagger-serve: swagger run

## docs: Alias untuk swagger
docs: swagger

.PHONY: help install build run run-dev test migrate db-reset clean mod-update lint format setup docker-build docker-run logs swagger swagger-serve docs
