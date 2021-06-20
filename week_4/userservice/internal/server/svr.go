package server

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Server interface {
	ListenAndServe(ctx context.Context, address string) error
	Shutdown(ctx context.Context) error
}

type UserServer struct {
	grpcServer *GrpcServer
	httpServer *HttpServer
	config     config
	ctx        context.Context
	cancelF    context.CancelFunc
}

type config struct {
	httpAddress string
	grpcAddress string
}

func NewUserServer(grpc *GrpcServer, http *HttpServer) *UserServer {
	return &UserServer{
		grpcServer: grpc,
		httpServer: http,
		config: config{
			httpAddress: ":80",
			grpcAddress: ":81",
		},
	}
}

func (server *UserServer) Start(ctx context.Context) error {
	newContext, cancelF := context.WithCancel(ctx)
	server.ctx = newContext
	server.cancelF = cancelF
	group, cctx := errgroup.WithContext(server.ctx)
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Printf("context done receive, exiting http server\n")
			server.httpServer.Shutdown(cctx)
		}()
		return server.httpServer.ListenAndServe(cctx, server.config.httpAddress)
	})
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Printf("context done receive, exiting grpc server\n")
			_ = server.grpcServer.Shutdown(cctx)

		}()
		return server.grpcServer.ListenAndServe(cctx, server.config.grpcAddress)
	})
	return group.Wait()
}

func (Server *UserServer) Shutdown() {
	Server.cancelF()
}
