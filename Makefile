include .env
export

export PROJECT_ROOT="."
run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/myapp/main.go

dcw:
	docker compose down

dcr: dcw
	docker compose up -d
#migrate-create seq=название миграции
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует параметр seq. make migrate-create seq=название-миграции"; \
		exit 1;\
	fi;\

	docker compose run --rm postgres-migrate create -ext sql -dir ./migrations -seq "$(seq)"\

migrate-up:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm postgres-migrate \
			-path ./migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
			up; \
	else \
		docker compose run --rm postgres-migrate \
			-path ./migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
			up "$(level)"; \
	fi

migrate-down:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm postgres-migrate \
			-path ./migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
			down; \
	else \
		docker compose run --rm postgres-migrate \
			-path ./migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
			down "$(level)"; \
	fi
