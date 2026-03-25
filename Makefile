MAIN_WEB = cmd/web/main.go
MAIN_CLI = cmd/cli/main.go
APP = board-backend
VERSION = 0.2.0

all: build-alpine

.PHONY: debug
debug:
	go run $(MAIN_WEB)

.PHONY: cli
cli:
	go run $(MAIN_CLI)

.PHONY: build-alpine
build-alpine:
	@echo "Building alpine image for backend server..."
	docker build -f Dockerfile.alpine -t $(APP):alpine-$(VERSION) .
	docker tag $(APP):alpine-$(VERSION) $(APP):alpine-latest

.PHONY: build-distroless
build-distroless:
	@echo "Building distroless image for backend server..."
	docker build -f Dockerfile.distroless -t $(APP):distroless-$(VERSION) .
	docker tag $(APP):distroless-$(VERSION) $(APP):distroless-latest

.PHONY: run-alpine
run-alpine:
	docker run -d -p 8080:8080 --name $(APP)-alpine $(APP):alpine-latest

.PHONY: run-distroless
run-distroless:
	docker run -d -p 8081:8080 --name $(APP)-distroless $(APP):distroless-latest

.PHONY: start-alpine
start-alpine:
	docker start $(APP)-alpine

.PHONY: start-distroless
start-distroless:
	docker start $(APP)-distroless

.PHONY: stop-alpine
stop-alpine:
	docker stop $(APP)-alpine

.PHONY: stop-distroless
stop-distroless:
	docker stop $(APP)-distroless
