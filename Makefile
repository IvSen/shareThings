APP_BIN = cmd/shareThings/build
DOCKER_COMPOSE_WEB ?= deployments/docker/web/docker-compose.yml

.PHONY: web_env
web_env: web_env_stop
	docker-compose -f ${DOCKER_COMPOSE_WEB} up -d

.PHONY: web_env_stop
web_env_stop:
	docker-compose -f ${DOCKER_COMPOSE_WEB} down


.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/shareThings/main.go

.PHONY: clean
clean:
	rm -rf ./app/build || true

.PHONY: swagger
swagger:
	swag init -g ./cmd/shareThings/main.go -o ./docs

.PHONY: migrate
migrate:
	$(APP_BIN) migrate -version $(version)

.PHONY: migrate.down
migrate.down:
	$(APP_BIN) migrate -seq down

.PHONY: migrate.up
migrate.up:
	$(APP_BIN) migrate -seq up