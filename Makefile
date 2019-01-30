docker: ## Run containers
	docker-compose up -d --build

docker-down: ## Shutdown containers
	docker-compose stop && docker-compose rm -f

account: ## Run account service
	go run ./cmd/account

account-migrate-up: ## Run migrations for account database
	go run ./cmd/account/migrate

account-migrate-drop: ## Drop account database
	go run ./cmd/account/migrate -action=drop

account-fake: ## Insert fake data into account database
	go run ./cmd/account/fakedata

help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
