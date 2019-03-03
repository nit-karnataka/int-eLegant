package main

import (
	proto "flok-server/auth-srv/authproto"
	"flok-server/auth-srv/crypto"
	"flok-server/auth-srv/handler"
	"flok-server/lib"
	"context"
	"log"
	"net"
	"time"

	"github.com/go-redis/redis"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	port     string
	dbName   string
	dbURL    string
	cacheURL string
	server   *grpc.Server
	listener *net.Listener
	store    *lib.Store
	hash     *crypto.Hash
}

func newApp(port, dbName, dbURL, cacheURL string) *app {
	return &app{
		port:     port,
		dbName:   dbName,
		dbURL:    dbURL,
		store:    &lib.Store{},
		hash:     &crypto.Hash{},
		cacheURL: cacheURL,
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

	client := redis.NewClient(&redis.Options{
		Addr:     a.cacheURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("Redis connected")

	proto.RegisterAuthServiceServer(a.server, &handler.AuthServiceHandler{
		Store:       a.store,
		Hash:        a.hash,
		CacheClient: client,
	})
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
	/* lis, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}

	a.server = grpc.NewServer()

	proto.RegisterHubServiceServer(a.server, &handler.HouseServiceHandler{})
	// Register reflection service on gRPC server.
	reflection.Register(a.server) */
	// log.Printf("DB %+v", db.GetMongoSession())
	log.Println("Starting server")
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
