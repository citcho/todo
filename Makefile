.PHONY: up down logs ps test test-integration migrate generate generate-key help
.DEFAULT_GOAL := help

up: ## Do docker compose up with live reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	docker compose exec api go test -v -race -shuffle=on ./...

test-integration: ## Execute integration tests
	docker compose exec api go test -v -tags=integration ./...

migrate: ## Execute migration
	docker compose exec -T api mysqldef -u citcho -p Secretp@ssw0rd -h db -P 3306 todo_db < ./_tools/mysql/schema.sql

generate: ## Generate codes
	docker compose exec api go generate ./...

generate-key: ## Generate key pair for JWT
	openssl genrsa -out ./internal/pkg/auth/cert/secret.pem 4096 && \
    openssl rsa -in ./internal/pkg/auth/cert/secret.pem -pubout -out ./internal/pkg/auth/cert/public.pem

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'