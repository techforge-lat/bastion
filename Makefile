# Environment variables for project
-include $(PWD)/.env

# Export all variable to sub-make
export

# Run app commands

build:
	@go build -o ./bin/restapi ./cmd/restapi

run: build
	clear
	@./bin/restapi

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
    awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

fix-imports:
	gci -w .

# End app commands

# Start database commands

install-migrate:
	@brew install golang-migrate

migration-create:
	@migrate create -ext sql -dir ./database/migrations $(name)

DATABASE_URL=$(BASTION_DATABASE_DRIVER)://$(BASTION_DATABASE_USER):$(BASTION_DATABASE_PASSWORD)@$(BASTION_DATABASE_HOST):$(BASTION_DATABASE_PORT)/$(BASTION_DATABASE_NAME)?sslmode=$(BASTION_DATABASE_SSLMODE)
migration-up:
	@echo $(DATABASE_URL)
	@migrate -source file://database/migrations -database $(DATABASE_URL)  up $(count)

migration-down:
	@migrate -source file://database/migrations -database $(DATABASE_URL) down $(count)

migration-db-drop:
	@migrate -source file://database/migrations -database $(DATABASE_URL) drop

# End database commands

install-lint:
	sudo curl -sSfL \
 https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh\
 | sudo sh -s -- -b /usr/local/bin v1.56.2

lint:
	golangci-lint run --fast

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.

