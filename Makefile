.PHONY: migrate-up migrate-down migrate-force

migrate-up:
	migrate -path migrations -database "$$DSN" up

migrate-down:
	migrate -path migrations -database "$$DSN" down

migrate-force:
	migrate -path migrations -database "$$DSN" force $(version)
