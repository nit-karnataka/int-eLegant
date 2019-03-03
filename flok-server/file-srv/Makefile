.PHONY: proto dev build

proto:
	protoc -I fileproto fileproto/*.proto --go_out=plugins=grpc:fileproto
	protoc-go-inject-tag -input=./fileproto/file.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build