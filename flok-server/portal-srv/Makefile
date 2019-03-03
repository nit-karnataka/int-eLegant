.PHONY: proto dev build

proto:
	protoc -I portalproto portalproto/*.proto --go_out=plugins=grpc:portalproto
	protoc-go-inject-tag -input=./portalproto/portal.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build