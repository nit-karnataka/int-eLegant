.PHONY: proto dev build

proto:
	protoc -I formproto formproto/*.proto --go_out=plugins=grpc:formproto
	protoc-go-inject-tag -input=./formproto/form.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build