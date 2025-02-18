# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

setup:
	@echo "Setting up the project..."
	@go mod tidy
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Go modules & dependencies installed successfully."

db-up:
	@echo "Starting PostgreSQL database in Docker..."
	@docker-compose --env-file .env up -d
	@echo "PostgreSQL database started."

db-down:
	@echo "Stopping PostgreSQL database container..."
	@docker-compose down
	@echo "PostgreSQL database stopped."

# Goose commands
migrate-up:
	@echo "Running database migrations..."
	@source .env && goose -dir database/migrations postgres "host=$${DB_HOST} port=$${DB_PORT} user=$${DB_USER} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" up
	@echo "Database migrations applied."
	
migrate-down:
	@echo "Reverting database migrations..."
	@source .env && goose -dir database/migrations postgres "host=$${DB_HOST} port=$${DB_PORT} user=$${DB_USER} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" down
	@echo "Database migrations reverted."

migrate-create:
	@echo "Creating new migration file..."
	@test -n "$(name)" || { echo "Error: 'name' argument is required. Usage: make migrate-create name=add_table"; exit 1; }
	@goose -dir database/migrations create "$(name)" sql
	@echo "Migration created: database/migrations/$(name).sql"

test:
	@echo "Running all tests..."
	@go test ./... -v
	@echo "All tests completed."

run:
	@echo "Starting the application..."
	@go run ./cmd/main.go
	@echo "Application started."

# Cleaning up Docker containers
clean:
	@echo "Cleaning up Docker containers..."
	@docker-compose down -v
	@echo "Cleanup complete."
