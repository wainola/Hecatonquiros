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
	listItems := []*list.List{
		&list.List{Id: "123", Content: "Content number 1"},
		&list.List{Id: "111", Content: "Content number 2"},
		&list.List{Id: "456", Content: "Content number 3"},
	}
	fmt.Println("req:", req.GetListId())

	return &list.ListResp{Items: listItems}, nil
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
