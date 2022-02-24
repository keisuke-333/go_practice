include .env
.PHONY: migrate_up
migrate_up:
	goose -dir ./db/migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

.PHONY: migrate_down
migrate_down:
	goose -dir ./db/migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down