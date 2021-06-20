//+build wireinject

package main

import (
	"github.com/google/wire"
	"wwbweibo.icu/userservice/internal/biz"
	"wwbweibo.icu/userservice/internal/data"
	"wwbweibo.icu/userservice/internal/server"
)

func InitServer() *server.UserServer {
	wire.Build(server.NewUserServer,
		server.NewGrpcServer,
		server.NewHttpServer,
		biz.NewUserService,
		data.NewInMemoUserDao)
	return &server.UserServer{}
}
