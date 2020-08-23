package main

import (
	"context"
	list "github/wainola/proglog/grpcV3/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := list.NewListServiceClient(conn)

	printList(client)
}

func printList(client list.ListServiceClient) {
	log.Print("Getting the fucking list item!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rr := &list.GetListReq{}
	listItem, err := client.GetList(ctx, rr)

	if err != nil {
		log.Fatalf("Error on getting item list")
	}
	log.Println("List item:::", listItem)
}
