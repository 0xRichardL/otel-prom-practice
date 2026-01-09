.PHONY: up down clean

export COMPOSE_PARALLEL_LIMIT=4

up:
	docker-compose up -d --build

down:
	docker-compose down

clean:
	docker-compose down -v --rmi all --remove-orphans