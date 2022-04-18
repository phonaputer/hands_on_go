compile:
	go build -o server_exe cmd/server.go

format:
	go fmt ./...

build: format compile

run:
	./server_exe
