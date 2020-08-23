package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	list "github/wainola/proglog/grpcV3/proto"

	"google.golang.org/grpc"
)

type listServer struct{}

func (l *listServer) GetList(ctx context.Context, req *list.GetListReq) (*list.ListResp, error) {
	fmt.Println("req:", req.GetListId())

	listItem := list.List{Id: "123", Content: "New Content"}

	return &list.ListResp{Items: []*list.List{&listItem}}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	srv := grpc.NewServer()
	list.RegisterListServiceServer(srv, &listServer{})
	reflection.Register(srv)

	fmt.Println("GRPC Server!")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
