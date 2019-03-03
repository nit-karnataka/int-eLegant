.PHONY: proto dev build

proto:
	protoc -I userproto userproto/user.proto --go_out=plugins=grpc:userproto
	protoc-go-inject-tag -input=./userproto/user.pb.go
dev:
	go build -o dev-server
	./dev-server

build:
	go build
