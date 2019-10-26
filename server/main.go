package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	user "grpc/proto"
	"log"
	"net"
)

type UserServer struct {
}

var users = make(map[uint64]*user.UserMessage)

func (u *UserServer) Set(_ context.Context, user *user.UserMessage) (*user.UserId, error) {
	if user.Id == nil {
		return nil, status.Error(codes.InvalidArgument, "userId not found")
	}
	users[user.Id.Value] = user
	return user.Id, nil
}

func (u *UserServer) Get(_ context.Context, userId *user.UserId) (*user.UserMessage, error) {
	_user := users[userId.Value]
	if _user == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return _user, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:4555")
	log.Println("Server starts")
	if err != nil {
		log.Fatal(err)
	}
	serv := grpc.NewServer()
	user.RegisterUserServer(serv, &UserServer{})
	err = serv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
