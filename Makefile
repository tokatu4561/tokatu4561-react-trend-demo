STRIPE_SECRET=
STRIPE_KEY=
DATABASE_URL=

## migrate: migration db data
migrate:
	@echo "migration db data..."
	@migrate -database ${DATABASE_URL} -path database/migrations up

