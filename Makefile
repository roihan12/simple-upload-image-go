#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up: ## Run migrations UP
migrate-up: 
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up
	docker compose up

migrate-down: ## Rollback migrations against non test DB
migrate-down: 
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migrate-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-create: 
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migration $(name)

shell-db: ## Enter to database console
shell-db: 
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres