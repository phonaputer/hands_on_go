BINARY=userapp

compile:
	go build -o $(BINARY) cmd/userapp.go

format:
	go fmt ./...

build: format compile

run_local: export USER_APPLICATION_CONFIG=deployments/local/user_app_config.json
run_local:
	./$(BINARY)

test:
	go test -cover ./...
