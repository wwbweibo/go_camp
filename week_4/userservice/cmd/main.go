package main

import "context"

//+build wireinject

func main() {
	server := InitServer()
	server.Start(context.Background())
}
