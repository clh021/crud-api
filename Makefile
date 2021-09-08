.PHONY: generate test serve

generate:
	@go mod tidy
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

gitTime=$(shell date +00%y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD)
# gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
build: generate
	@go build -ldflags "-X main.build=${gitTime}.${gitCID}" -o "bin/crud-api"
	@echo "[OK] App binary was created!"

run:
	@./bin/crud-api

test: 
	go test -v ./...

serve: build run
