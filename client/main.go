package main

import (
	"context"
	"google.golang.org/grpc"
	user "grpc/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:4555", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := user.NewUserClient(conn)
	resp, err := client.Set(context.Background(), &user.UserMessage{
		Name: "John",
		Age:  0,
		Id: &user.UserId{
			Value: 10,
		},
		ChildCount: &user.NatNumber{Value: 2},
		CatsCount:  &user.NatNumber{Value: 7},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	respGet, err := client.Get(context.Background(), &user.UserId{Value: 10})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(respGet)
}
