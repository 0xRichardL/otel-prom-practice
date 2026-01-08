.PHONY: up down

export COMPOSE_PARALLEL_LIMIT=4

up:
	docker-compose up -d --build

down:
	docker-compose down