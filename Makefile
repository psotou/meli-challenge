up:
	docker compose up --build
down:
	docker compose down
volumes-down:
	docker compose down -v
bash:
	docker compose exec mysqldb /bin/bash
db:
	docker compose exec mysqldb mysql -u root -p
logs:
	docker compose logs
