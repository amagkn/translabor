export

DB_MIGRATE_URL = postgres://postgres:postgres@localhost:5432/translabor?sslmode=disable
MIGRATE_PATH = ./migration/postgres

mod:
	go mod tidy

dev: mod
	go run ./cmd/app

start:
	go build -o ./bin/app ./cmd/app && ./bin/app

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $(name)

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all