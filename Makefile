.PHONY: up upd help gen
.DEFAULT_GOAL := help

up: ## Do docker compose up with live reload
	docker compose up --build

upd: ## Do docker compose up demon with live reload
	docker compose up -d --build

down: ## Do docker compose down
	docker compose down

gen:
	gqlgen generate

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

