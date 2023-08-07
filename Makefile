include .env
export

MYSQL_DSN ?= $(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

local-db:
	@ docker-compose up -d


	@ until mysql --host=$(AUTH_MYSQL_HOST) --port=$(AUTH_MYSQL_PORT) --user=$(AUTH_MYSQL_USERNAME) -p$(AUTH_MYSQL_PASSWORD) --protocol=tcp -e 'SELECT 1' >/dev/null 2>&1 && exit 0; do \
	  >&2 echo "MySQL is unavailable - sleeping"; \
	  sleep 5 ; \
	done


	@ echo "MySQL is up and running!"


migrate-setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing migrate command..."; go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate; fi


migrate-up: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations up $(N)


migrate-down: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations down $(N)


migrate-to-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations goto $(V)


drop-db: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations drop


force-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations force $(V)


migration-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations version


build:
	@ go build main.go

run: build
	$(env) go run .