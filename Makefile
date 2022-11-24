build:
	docker compose build
start:
	docker compose up -d
stop:
	docker compose down
restart:
	make stop
	make build
	make start
logs:
	docker compose logs -f