include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
	.PHONY: confirm confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/dev: run the cmd/api application
.PHONY: run/dev
run:
	go run ./cmd/
compile:
	cd cmd && go run github.com/99designs/gqlgen compile



## db/psql: connect to the database using psql
.PHONY: db/mysql
db/mysql:
	mysql ${DB_DSN}
	
## db/migrations/new name=$1: create a new database migration
.PHONY: db/new
 db/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}
	
## db/migrations/up: apply all up database migrations
.PHONY: db/up
db/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database mysql://"${DEV_DB_DSN}" up

.PHONY: db/down
db/down:
	@echo 'Running down migrations...'
	migrate -path ./migrations -database mysql://"${DEV_DB_DSN}" down -all

db/force:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database mysql://"${DEV_DB_DSN}"  force ${v}


db/up/prod:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database mysql://"${PROD_DB_USER}:${PROD_DB_PASSWORD}@tcp(${PROD_DB_HOST}:${PROD_DB_PORT})/prod_db?allowCleartextPasswords=true" up

db/go:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database mysql://"${DEV_DB_DSN}"  goto ${v}

db/version:
	@echo "Current Version:"
	@migrate -path ./migrations -database mysql://"${DEV_DB_DSN}" version


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...' go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
## vendor: tidy and vendor dependencies
.PHONY: vendor
 vendor:
	@echo 'Tidying and verifying module dependencies...' go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

build:
	@echo 'Building optimized docker img...'
	make build/api
	docker build -t web-server .


build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./deployment/dev/current_os ./cmd/
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./deployment/prod/linux_amd64 ./cmd/


# ==================================================================================== #
# Deploy
# ==================================================================================== #

.PHONY: deploy/prod
deploy/prod:
	@echo 'Deploying to docker container'
	docker-compose -f docker-compose.prod.yml up --build -d

PHONY: deploy/dev
deploy/dev:
	@echo 'Deploying to docker container'
	docker-compose -f docker-compose.dev.yml up --build -d


# ==================================================================================== #
# Tests
# ==================================================================================== #

.PHONY: test/unit
test/unit:
	@echo 'Running all unit tests'
	go test ./... --short

.PHONY: test/unit/v
test/unit/v:
	@echo 'Running all unit tests'
	go test ./... --v --short


PHONY: test/int
test/int:
	@echo 'Running all test (+ integration), this may take a while'
	go test ./... -parallel 4

PHONY: test/int/v
test/int/v:
	@echo 'Running all test (+ integration), this may take a while'
	go test ./... -v
