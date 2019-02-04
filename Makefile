docker: ## Run containers
	docker-compose up -d --build

docker-down: ## Shutdown containers
	docker-compose stop && docker-compose rm -f

api: ## Run api service
	go run ./cmd/api

user: ## Run user service
	go run ./cmd/user

user-migrate-up: ## Run migrations for user database
	go run ./cmd/user/migrate

user-migrate-drop: ## Drop user database
	go run ./cmd/user/migrate -action=drop

user-fake: ## Insert fake data into user database
	go run ./cmd/user/fakedata

scream: ## Run scream service
	go run ./cmd/scream

scream-migrate-up: ## Run migrations for scream database
	go run ./cmd/scream/migrate

scream-migrate-drop: ## Drop scream database
	go run ./cmd/scream/migrate -action=drop

scream-fake: ## Insert fake data into scream database
	go run ./cmd/scream/fakedata

help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
