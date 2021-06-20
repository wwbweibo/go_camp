// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"wwbweibo.icu/userservice/internal/biz"
	"wwbweibo.icu/userservice/internal/data"
	"wwbweibo.icu/userservice/internal/server"
)

// Injectors from wire.go:

func InitServer() *server.UserServer {
	userDaoIntf := data.NewInMemoUserDao()
	userService := biz.NewUserService(userDaoIntf)
	grpcServer := server.NewGrpcServer(userService)
	httpServer := server.NewHttpServer(userService)
	userServer := server.NewUserServer(grpcServer, httpServer)
	return userServer
}