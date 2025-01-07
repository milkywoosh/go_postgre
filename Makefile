# Container name
CONTAINER_NAME = pg17v1

# Default target
.PHONY: all
all: help


# Display help
.PHONY: help
help:
	@echo "Makefile for managing PostgreSQL Docker container $(CONTAINER_NAME):"
	@echo "  make start   - Start the container"
	@echo "  make stop    - Stop the container"
	@echo "  make restart - Restart the container"
	@echo "  make status  - Check container status"
	@echo "  make help    - Show this help message"

.PHONY: start
start:
	@echo "Starting container $(CONTAINER_NAME)..."
	docker start $(CONTAINER_NAME)

.PHONY: stop
stop:
	@echo "Stop container $(CONTAINER_NAME)"
	docker stop $(CONTAINER_NAME)

.PHONY: status
status:
	@echo "Checking status of container $(CONTAINER_NAME)..."
	docker ps -a --filter "name=$(CONTAINER_NAME)"

.PHONY: toko_buku_users
toko_buku_users:
	docker exec -it pg17v1 psql -U postgres -d toko_buku_online_nextjs select * from users

.PHONY: check_user_roles
check_user_roles:
	docker exec -it pg17v1 psql -U postgres -d toko_buku_online_nextjs -c "select u.username, r.role_name from users u \
	left join user_roles ur on u.id = ur.user_id \
	left join roles r on r.id = ur.role_id \
	where username='luke'"

.PHONY: run_any_query
run_any_query:
	docker exec -it pg17v1 psql -U postgres -d toko_buku_online_nextjs -c "\
	select u.username \
	from users u \
	where u.username='luke'"


