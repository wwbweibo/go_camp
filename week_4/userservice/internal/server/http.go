package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	pb "wwbweibo.icu/userservice/api/v1"
	"wwbweibo.icu/userservice/internal/biz"
	"wwbweibo.icu/userservice/internal/data"
	"wwbweibo.icu/userservice/internal/model"
)

type HttpServer struct {
	userService *biz.UserService
	router      *gin.Engine
	svr         *http.Server
}

func NewHttpServer(userService *biz.UserService) *HttpServer {
	server := &HttpServer{
		userService: userService,
		router:      gin.Default(),
		svr:         &http.Server{},
	}
	server.router.POST("/userservice/api/v1/register", server.Register)
	server.router.POST("/userservice/api/v1/login", server.Login)
	server.svr.Handler = server.router
	return server
}

func (server *HttpServer) Register(ctx *gin.Context) {
	userName := ""
	password := ""
	email := ""
	if userName, ok := ctx.GetPostForm("username"); !ok || userName == "" {
		ctx.JSON(400, pb.Response{
			Code:    400,
			Message: "user name can not be ignored or empty",
		})
		return
	}
	if password, ok := ctx.GetPostForm("password"); !ok || password == "" {
		ctx.JSON(400, pb.Response{
			Code:    400,
			Message: "password can not be ignored or empty",
		})
		return
	}
	if email, ok := ctx.GetPostForm("password"); !ok || email == "" {
		ctx.JSON(400, pb.Response{
			Code:    400,
			Message: "email can not be ignored or empty",
		})
		return
	}

	result, err := server.userService.Register(ctx.Request.Context(), model.User{userName, password, email})
	if err != nil {
		if errors.Is(err, data.ErrAlreadyExist) {
			ctx.JSON(400, pb.Response{
				Code:    400,
				Message: fmt.Sprintf("user with user name %s is already exist", userName),
			})
		} else {
			fmt.Printf("GrpcServer error, error while create user, %s", err)
			ctx.JSON(500, pb.Response{
				Code:    500,
				Message: "internal server error",
			})
		}
		return
	}
	if result {
		ctx.JSON(200, pb.Response{
			Code:    200,
			Message: "success",
			Data:    true,
		})
	} else {
		ctx.JSON(200, pb.Response{
			Code:    200,
			Message: "failed",
			Data:    false,
		})
	}
}

func (server *HttpServer) Login(ctx *gin.Context) {
	userName := ""
	password := ""
	if userName, ok := ctx.GetPostForm("username"); !ok || userName == "" {
		ctx.JSON(400, pb.Response{
			Code:    400,
			Message: "user name can not be ignored or empty",
		})
		return
	}
	if password, ok := ctx.GetPostForm("password"); !ok || password == "" {
		ctx.JSON(400, pb.Response{
			Code:    400,
			Message: "password can not be ignored or empty",
		})
		return
	}

	result, err := server.userService.Login(ctx.Request.Context(), userName, password)
	if err != nil {
		if errors.Is(err, data.ErrNotExist) {
			ctx.JSON(200, pb.Response{
				Code:    200,
				Message: "could find user",
			})
		} else {
			fmt.Printf("GrpcServer error, error while login user, %s", err)
			ctx.JSON(500, pb.Response{
				Code:    500,
				Message: "internal server error",
			})
		}
		return
	}
	if result {
		ctx.JSON(200, pb.Response{
			Code:    200,
			Message: "success",
			Data:    true,
		})
	} else {
		ctx.JSON(200, pb.Response{
			Code:    200,
			Message: "password error",
		})
	}
}

func (server *HttpServer) ListenAndServe(ctx context.Context, address string) error {
	server.svr.Addr = address
	return server.svr.ListenAndServe()
}
func (server *HttpServer) Shutdown(ctx context.Context) error {
	return server.svr.Shutdown(ctx)
}
