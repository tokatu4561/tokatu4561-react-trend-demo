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

down-migrate:
	@echo "migration db data..."
	@migrate -database ${DATABASE_URL} -path backend/database/migrations/ down

seed:
	@echo "seeding db data..."
	@go run database/seeder/seed.go

down-all:
	@docker-compose down --rmi all --volumes --remove-orphans

yarn-dev:
	@cd ./frontend && yarn dev

run-api:
	@cd ./backend&& go run main.go route.go helpers.go handler.go


# テーブル作成コマンド
# migrate create -ext sql -dir db/migrations -seq create_users
# apiサーバー