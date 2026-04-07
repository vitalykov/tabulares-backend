MAIN_WEB = cmd/web/main.go
MAIN_CLI = cmd/cli/main.go
APP = tabulares-backend
VERSION = 0.3.0

all: build

.PHONY: debug
debug:
	go run $(MAIN_WEB)

.PHONY: cli
cli:
	go run $(MAIN_CLI)

.PHONY: build
build-alpine:
	@echo "Building alpine image for backend server..."
	docker build -f Dockerfile -t $(APP)-$(VERSION) .
	docker tag $(APP):$(VERSION) $(APP):latest

.PHONY: run
run:
	docker run -it --rm -p 8080:8080 --name $(APP) $(APP):latest

.PHONY: start
start:
	docker compose up

.PHONY: stop
stop:
	docker compose down
