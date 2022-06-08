compile:
	go build -o server_exe cmd/server.go

format:
	go fmt ./...

build: format compile

run: export HOG_CONFIG_PATH=./configs/local_config.json
run:
	./server_exe
