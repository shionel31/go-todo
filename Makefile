docker-up-%:
	@docker-compose -f docker-compose.${@:docker-up-%=%}.yml up

docker-down-%:
	@docker-compose -f docker-compose.${@:docker-down-%=%}.yml down

migrate-%:
	go run ./cmd/${@:migrate-%=%}/migrate/migrate.go

.PHONY: migrate
migrate:
	make migrate-todo
