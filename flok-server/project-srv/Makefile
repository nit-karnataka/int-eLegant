.PHONY: proto dev build

proto:
	protoc -I projectproto projectproto/*.proto --go_out=plugins=grpc:projectproto
	protoc-go-inject-tag -input=./projectproto/project.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build