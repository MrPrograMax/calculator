APP_NAME=calculator

.PHONY: build
build:
	docker compose build

.PHONY: run
run:
	docker compose up -d

.PHONY: stop
stop:
	docker compose down

.PHONY: logs
logs:
	docker compose logs -f

.PHONT: test
test:
	go test -cover ./...

.PHONT: test-total
test-total:
	go tool cover -func=coverage.out | grep total
