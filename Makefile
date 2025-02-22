.PHONY: migrate-up migrate-down
POSTGRES_URL=postgres://viet:123@localhost:5432/mpc_key?sslmode=disable
migrate-up:
	goose -dir db/migrations postgres "$(POSTGRES_URL)" up

migrate-down:
	goose -dir db/migrations postgres "$(POSTGRES_URL)" down
run:
	go run main.go