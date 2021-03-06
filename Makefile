BIN = $(CURDIR)/bin
PROJECT = habits
MAIN = cmd/habits/main.go
ADDR = localhost:3000

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECT).pid

MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally.
install: 
	GO111MODULE=on go get ./...

## start: Start in development mode. Auto-starts when code changes.
start: 
	@bash -c "trap 'make stop' EXIT; $(MAKE) build start-server watch run='make stop-server build start-server'"

## stop: Stop development mode.
stop: stop-server

start-server:
	@echo "  >  $(PROJECT) has been started"
	@$(BIN)/$(PROJECT) serve --listen-addr :3000 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@echo "  >  stopping $(PROJECT)"
	@touch $(PID)
	@kill `cat $(PID)` 2> /dev/null || true
	@rm $(PID)

restart-server: stop-server start-server

## build: Build the application
build: build-js build-api	

build-js:
	npm run build -- --mode development

build-js-production:
	npm run build

build-api:
	GO111MODULE=on go build -o $(BIN)/$(PROJECT) $(MAIN)

build-api-production:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $(BIN)/$(PROJECT) $(MAIN)

## test: Test the application
test:
	GO111MODULE=on go vet ./...
	GO111MODULE=on go test ./... -covermode=count -coverprofile=c.out ./...

## clean: Clean build files. Runs `go clean` internally.
clean:
	GO111MODULE=on go clean ./...

## fmt: Runs `go fmt` internally.
fmt:
	GO111MODULE=on go fmt ./...

## compile: Clean and then compile for production
compile: clean install build

## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
watch:
	yolo -i . -e node_modules -e bin -e web/dist -c "$(run)"

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo