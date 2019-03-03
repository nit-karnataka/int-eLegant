.PHONY: proto dev build

proto:
	protoc -I chatproto chatproto/*.proto --go_out=plugins=grpc:chatproto
	protoc-go-inject-tag -input=./chatproto/chat.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build