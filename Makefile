STRIPE_SECRET=
STRIPE_KEY=
DATABASE_URL='postgres://postgres:password@localhost:5432/practice?sslmode=disable'
# DSN="host=postgres port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"

up:
	@echo "build docker conteiner"
	@docker compose up -d --build

## migrate: migration db data
migrate:
	@echo "migration db data..."
	@migrate -database ${DATABASE_URL} -path backend/database/migrations/ up

down-all:
	@docker-compose down --rmi all --volumes --remove-orphans