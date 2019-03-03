.PHONY: proto dev build

proto:
	protoc -I authproto authproto/auth.proto --go_out=plugins=grpc:authproto
	protoc-go-inject-tag -input=./authproto/auth.pb.go
dev:
	go build -o dev-server
	./dev-server

build:
	go build