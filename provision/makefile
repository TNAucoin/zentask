.PHONY : rebuild-all down up reset-auth-service-db rebuild-auth-service


rebuild-all:
	@echo "Rebuilding..."
	docker compose down
	docker compose up -d --build
	@echo "Done."

down:
	@echo "Stopping..."
	docker compose down
	@echo "Done."

up:
	@echo "Starting..."
	docker compose up -d
	@echo "Done."

rebuild-auth-service:
	@echo "Rebuilding auth service..."
	docker compose up -d --build authentication-service
	@echo "Done."

reset-auth-service-db:
	@echo "Resetting auth service db..."
	docker compose down 
	docker volume rm auth_service_db_data
	docker compose up -d auth-service-db
	sleep 10

	docker exec -it $$(docker-compose ps -q) psql -U postgres -c "DROP TABLE IF EXISTS users;"
	docker exec -it $$(docker-compose ps -q) psql -U postgres -c "CREATE TABLE users();"

	@echo "Done."

	