include .env
.PHONY: migrate_up
migrate_up:
	goose -dir ./db/migrations $(DB_DRIVER) $(DB_DNS) up

.PHONY: migrate_down
migrate_down:
	goose -dir ./db/migrations $(DB_DRIVER) $(DB_DNS) down