# Default compose file path - can be overridden from command line
COMPOSE_FILE ?= docker/docker-compose.yml

.PHONY: up down ps logs clean

# Start all services
up:
	docker compose -f $(COMPOSE_FILE) up -d --remove-orphans

# Stop all services
down:
	docker compose -f $(COMPOSE_FILE) down

# Show running containers
ps:
	docker compose -f $(COMPOSE_FILE) ps

# Show logs from all services
logs:
	docker compose -f $(COMPOSE_FILE) logs -f

# Clean up containers, volumes, and orphaned containers
clean:
	docker compose -f $(COMPOSE_FILE) down -v --remove-orphans