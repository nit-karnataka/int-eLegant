package main

import (
	"context"
	"flok-server/lib"
	"flok-server/meeting-srv/handler"
	proto "flok-server/meeting-srv/meetingproto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	port     string
	dbName   string
	dbURL    string
	store    *lib.Store
	server   *grpc.Server
	listener *net.Listener
}

func newApp(port, dbName, dbURL string) *app {
	return &app{
		port:   port,
		dbName: dbName,
		dbURL:  dbURL,
		store:  &lib.Store{},
	}
}

func (a *app) init() error {
	err := a.store.Connect(a.dbURL, "", "", a.dbName)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}

	a.listener = &lis

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	h := &handler.MeetingServiceHandler{
		Store: a.store,
	}

	proto.RegisterMeetingServiceServer(a.server, h)
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
	log.Println("Starting meeting server")
	if err := a.server.Serve(*a.listener); err != nil {
		return err
	}

	return nil
}

func (a *app) close() {
	log.Println("Stopping server")
	a.server.GracefulStop()

	a.store.Close()
}
