.PHONY: generate

# gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
preflags:
	@gitTime=$(date +00%y%m%d%H%M%S)
	@gitCID=`git rev-parse HEAD`

generate:
	@go mod tidy
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build: generate preflags
	@go build -ldflags "-X main.build=${gitTime}.${gitCID}" -o "bin/crud-api"
	@echo "[OK] App binary was created!"

run:
	@./bin/crud-api

test: build run
