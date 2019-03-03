.PHONY: proto dev build

proto:
	protoc -I meetingproto meetingproto/*.proto --go_out=plugins=grpc:meetingproto
	protoc-go-inject-tag -input=./meetingproto/meeting.pb.go

dev:
	go build -o dev-server
	./dev-server

build:
	go build