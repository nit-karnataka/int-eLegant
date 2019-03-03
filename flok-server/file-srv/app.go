package main

import (
	"context"
	proto "flok-server/file-srv/fileproto"
	"flok-server/file-srv/handler"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	port     string
	server   *grpc.Server
	listener *net.Listener
}

func newApp(port string) *app {
	return &app{
		port: port,
	}
}

func (a *app) init() error {
	lis, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}

	a.listener = &lis

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	h := &handler.FileServiceHandler{}

	proto.RegisterFileServiceServer(a.server, h)
	// Register reflection service on gRPC server.
	reflection.Register(a.server)

	return nil
}

func unaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	//logging
	log.Printf("request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func (a *app) listen() error {
	log.Println("Starting file server")
	if err := a.server.Serve(*a.listener); err != nil {
		return err
	}

	return nil
}

func (a *app) close() {
	log.Println("Stopping server")
	a.server.GracefulStop()
}
