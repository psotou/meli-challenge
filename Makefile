up:
	docker compose up --build
down:
	docker compose down
volumes-down:
	docker compose down -v
db:
	docker compose exec mysqldb mysql -u root -p
bash:
	docker compose exec mysqldb /bin/bash
logs:
	docker compose logs
tests:
	go clean -testcache && go test -v -cover ./...
seed-data:
	cd scripts/sh-script/ && ./post-data.sh
